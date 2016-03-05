// Package stratnet implementes spec.Network to provide functionality to create
// and optimize strategies.
package stratnet

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/strategy"
)

const (
	ObjectTypeStratNet spec.ObjectType = "strat-net"
)

type Config struct {
	// Dependencies.
	Log     spec.Log
	Storage spec.Storage

	PatNet spec.Network

	// Settings.

	// Actions represents a list of action items, that are object types. These
	// are used to find the best performing combination by ordering them in a
	// certain way. Such a ordered list of actions is called a strategy
	// maintained by a neuron. The best strategy is represented by the highest
	// score.
	Actions []spec.ObjectType

	// MaxElements representes the maximum number of elements to fetch from a
	// list within one call. This prevents fetching too much data at once.
	MaxElements int
}

// DefaultConfig provides a default configuration to create a new strategy
// network by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		PatNet: nil,

		// Settings.
		Actions:     []spec.ObjectType{},
		MaxElements: 10,
	}

	return newConfig
}

// NewStratNet returns a new configured strategy network.
//
// The network makes use of different key namespaces, that can be anything of
// the following. Note that these key namespaces might be reused in other
// networks as well.
//
//     strategy:successes
//
//         Holds the weighted list of strategy IDs ordered by most successes.
//
//         ID1:score1,ID2:score2,...
//
//     strategy:<strategyID>
//
//         Holds the action sequence of a strategy.
//
//         action1,action2,...
//
//     context:<contextID>
//
//         Holds the list of strategy IDs associated with the given context.
//
//         strategyID1,strategyID2,...
//
func NewStratNet(config Config) (spec.Network, error) {
	newNetwork := &stratNet{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeStratNet,
	}

	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

type stratNet struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (sn *stratNet) Boot() {
	sn.Mutex.Lock()
	defer sn.Mutex.Unlock()

	if sn.Booted {
		return
	}
	sn.Booted = true

	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call Boot")
}

func (sn *stratNet) GetBestStrategy(imp spec.Impulse) (spec.Strategy, error) {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call GetBestStrategy")

	successes, err := sn.GetMostSuccesses(imp)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(successes) == 0 {
		return nil, maskAny(noStrategyError)
	}

	strategy, err := sn.GetStrategyBySuccesses(imp, successes)
	if err != nil {
		return nil, maskAny(err)
	}

	return strategy, nil
}

func (sn *stratNet) GetMostSuccesses(imp spec.Impulse) ([]string, error) {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call GetHighestScore")

	ctx := imp.GetCtx(sn)
	successes, err := sn.Storage.GetHighestScoredElements(ctx.GetKey("strategy:successes"), sn.MaxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	return successes, nil
}

func (sn *stratNet) GetStrategyByID(imp spec.Impulse, ID spec.ObjectID) (spec.Strategy, error) {
	ctx := imp.GetCtx(sn)
	value, err := sn.Storage.Get(ctx.GetKey("strategy:%s", ID))
	if err != nil {
		return nil, maskAny(err)
	}

	var newActions []spec.ObjectType
	for _, a := range strings.Split(value, ",") {
		newActions = append(newActions, spec.ObjectType(a))
	}
	if newActions == nil {
		return nil, maskAny(invalidStrategyError)
	}

	newConfig := strategy.DefaultConfig()
	newConfig.Actions = newActions
	newConfig.ID = ID
	newStrategy, err := strategy.NewStrategy(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newStrategy, nil
}

func (sn *stratNet) GetStrategyBySuccesses(imp spec.Impulse, successes []string) (spec.Strategy, error) {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call GetStrategyByScore")

	// Find all strategy IDs having the highest score.
	var highScore int
	var mapped map[int][]string
	for i, e := range successes {
		if i%2 == 0 {
			// We are only interested in the scores and the list has the following
			// scheme. So we skip the elements, that are indexed with an even number.
			//
			//     element1,score1,element2,score2,...
			//
			continue
		}

		n, err := strconv.Atoi(e)
		if err != nil {
			return nil, maskAny(err)
		}
		if n > highScore {
			highScore = n
		}

		if _, ok := mapped[n]; !ok {
			mapped[n] = []string{}
		}

		mapped[n] = append(mapped[n], successes[i-1])
	}
	highestScored := mapped[highScore]
	i := rand.Intn(len(highestScored))
	ID := highestScored[i]

	newStrategy, err := sn.GetStrategyByID(imp, spec.ObjectID(ID))
	if err != nil {
		return nil, maskAny(err)
	}

	return newStrategy, nil
}

func (sn *stratNet) NewStrategy(imp spec.Impulse) (spec.Strategy, error) {
	ctx := imp.GetCtx(sn)

	newConfig := strategy.Config{
		Actions: sn.Actions,
	}

	var err error
	var newStrategy spec.Strategy
	for i := 0; i < 3; i++ {
		newStrategy, err = strategy.NewStrategy(newConfig)
		if err != nil {
			return nil, maskAny(err)
		}

		// Check if strategy already exists.
		//
		// TODO this needs to be improved. There are already ideas. See
		// https://github.com/xh3b4sd/anna/issues/79.
		_, err := sn.Storage.Get(ctx.GetKey("strategy:%s", newStrategy.String()))
		if err != nil {
			return nil, maskAny(err)
		}
	}

	if newStrategy == nil {
		// After 3 tries there is still no valid strategy. We already created all
		// combinations of the available actions. We cannot find any new strategy.
		// The given actions are not sufficient and lead to absolutely nothing.
		// Return an error for this serious problem.
		//
		// TODO that there is no combination actually might not be true. The more
		// combinations we already have, the more increases the probability to
		// create the same strategy three times in a row. We need to improve
		// strategy creation. A random order is a good first step, but is not
		// sufficient when we are going to reach the other end of the propability.
		// For some ideas see https://github.com/xh3b4sd/anna/issues/80.
		return nil, maskAnyf(combinationLimitError, "no more strategies left")
	}

	err = sn.StoreStrategy(imp, newStrategy)
	if err != nil {
		return nil, maskAny(err)
	}

	return newStrategy, nil
}

func (sn *stratNet) Shutdown() {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call Shutdown")
}

func (sn *stratNet) StoreStrategy(imp spec.Impulse, newStrategy spec.Strategy) error {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call StoreStrategy")

	ctx := imp.GetCtx(sn)
	err := sn.Storage.Set(ctx.GetKey("strategy:%s", newStrategy.GetID()), newStrategy.String())
	if err != nil {
		return maskAny(err)
	}
	err = sn.Storage.Set(ctx.GetKey("context:%s", ctx.GetID()), string(newStrategy.GetID()))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (sn *stratNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call Trigger")

	newStrategy, err := sn.GetBestStrategy(imp)
	if IsNoStrategy(err) {
		newStrategy, err = sn.NewStrategy(imp)
		if err != nil {
			return nil, maskAny(err)
		}
	} else if err != nil {
		return nil, maskAny(err)
	}

	// TODO add strategy to impulse
	fmt.Printf("%#v\n", newStrategy)

	return imp, nil
}
