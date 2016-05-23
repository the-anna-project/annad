// Package stratnet implementes spec.Network to provide functionality to create
// and optimize strategies.
package stratnet

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/strategy"
)

const (
	// ObjectTypeStratNet represents the object type of the strategy network
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeStratNet spec.ObjectType = "strat-net"
)

// Config represents the configuration used to create a new strategy network
// object.
type Config struct {
	// Dependencies.
	Log     spec.Log
	Storage spec.Storage

	PatNet spec.Network

	// Settings.

	// CLGNames represents a list of action items, that are CLG names. These are
	// used to find the best performing combination by ordering them in a certain
	// way. Such a ordered list of CLG names is called a strategy. The best
	// strategy is represented by the highest score.
	CLGNames []string

	// MaxElements representes the maximum number of elements to fetch from a
	// list within one call. This prevents fetching too much data at once.
	MaxElements int
}

// DefaultConfig provides a default configuration to create a new strategy
// network object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

		PatNet: nil,

		// Settings.
		CLGNames:    []string{},
		MaxElements: 10,
	}

	return newConfig
}

// NewStratNet creates a new configured strategy network object.
//
// The network makes use of different key namespaces, that can be anything of
// the following. Note that these key namespaces might be reused in other
// networks as well.
//
// Note the requestor is used to differentiate generated strategies for
// different purposes.
//
//     strategy:<requestor>:max:success
//
//         Holds the weighted list of strategy IDs ordered from least to most
//         successes.
//
//         ID1,score1,ID2,score2,...
//
//     strategy:<requestor>:data:<strategyID>
//
//         Holds the strategy's data.
//
//         key: value
//         key: value
//         ...
//
//     strategy:<requestor>:actions:<list>
//
//         Holds the comma separated representation of a strategy's generated
//         action sequence.
//
//     strategy:<requestor>:min:score
//
//         Holds the number of the minimal required score a strategy needs to
//         fulfil to be evaluated as sufficient.
//
//         number
//
func NewStratNet(config Config) (spec.Network, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newNetwork := &stratNet{
		Config:       config,
		BootOnce:     sync.Once{},
		ID:           newID,
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         ObjectTypeStratNet,
	}

	newNetwork.Log.Register(newNetwork.GetType())

	return newNetwork, nil
}

type stratNet struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (sn *stratNet) Boot() {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call Boot")

	sn.BootOnce.Do(func() {
	})
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
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call GetMostSuccesses")

	key := sn.key("strategy:%s:max:success", imp.GetRequestor())
	successes, err := sn.Storage.GetHighestScoredElements(key, sn.MaxElements)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(successes) < 2 {
		return nil, nil
	}
	highestScore, err := strconv.Atoi(successes[1])
	if err != nil {
		return nil, maskAny(err)
	}

	// Here it is likely to return the same weak strategies over and over again.
	// We need a dynamic adjustment of the successes that are accepted or denied.
	// That way we are able to create new strategies in case there are no
	// sufficient strategies found so far. Thus we compare the minimal required
	// score a strategy needs to fulfil to be evaluated as sufficient.
	minScore, err := sn.Storage.Get(sn.key("min:score"))
	if err != nil {
		return nil, maskAny(err)
	}
	requiredScore, err := strconv.Atoi(minScore)
	if err != nil {
		return nil, maskAny(err)
	}
	if highestScore < requiredScore {
		return nil, nil
	}

	return successes, nil
}

func (sn *stratNet) GetStrategyByID(imp spec.Impulse, ID spec.ObjectID) (spec.Strategy, error) {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call GetStrategyByID")

	key := sn.key("strategy:%s:data:%s", imp.GetRequestor(), ID)
	value, err := sn.Storage.Get(key)
	if err != nil {
		return nil, maskAny(err)
	}

	if value == "" {
		return nil, maskAny(strategyNotFoundError)
	}

	newStrategy := strategy.NewEmptyStrategy()
	err = json.Unmarshal([]byte(value), newStrategy)
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
		if i%2 != 0 {
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
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call NewStrategy")

	newConfig := strategy.Config{
		CLGNames:  imp.GetCLGNames(),
		Requestor: imp.GetRequestor(),
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
		key := sn.key("strategy:%s:actions:%s", imp.GetRequestor(), strings.Join(newStrategy.GetCLGNames(), ","))
		_, err := sn.Storage.Get(key)
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

	sn.ShutdownOnce.Do(func() {
	})
}

func (sn *stratNet) StoreStrategy(imp spec.Impulse, strat spec.Strategy) error {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call StoreStrategy")

	key := sn.key("strategy:%s:data:%s", imp.GetRequestor(), strat.GetID())
	raw, err := json.Marshal(strat)
	if err != nil {
		return maskAny(err)
	}
	err = sn.Storage.Set(key, string(raw))
	if err != nil {
		return maskAny(err)
	}
	key = sn.key("strategy:%s:actions:%s", imp.GetRequestor(), strings.Join(strat.GetCLGNames(), ","))
	err = sn.Storage.Set(key, string(strat.GetID()))
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (sn *stratNet) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	sn.Log.WithTags(spec.Tags{L: "D", O: sn, T: nil, V: 13}, "call Trigger")

	newStrategy, err := sn.GetBestStrategy(imp)
	if IsNoStrategy(err) {
		// No sufficient strategy could be found. Lets create a new one.
		newStrategy, err = sn.NewStrategy(imp)
		if err != nil {
			return nil, maskAny(err)
		}
	} else if err != nil {
		return nil, maskAny(err)
	}

	imp.SetStrategyByRequestor(imp.GetRequestor(), newStrategy)

	return imp, nil
}
