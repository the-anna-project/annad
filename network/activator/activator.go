package activator

import (
	"strings"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the activator object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "activator"
)

// Config represents the configuration used to create a new activator object.
type Config struct {
	// Dependencies.
	Log               spec.Log
	FactoryCollection spec.FactoryCollection
	StorageCollection spec.StorageCollection
}

// DefaultConfig provides a default configuration to create a new activator
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log:               log.New(log.DefaultConfig()),
		FactoryCollection: factory.MustNewCollection(),
		StorageCollection: storage.MustNewCollection(),
	}

	return newConfig
}

// New creates a new configured activator object.
func New(config Config) (spec.Activator, error) {
	newActivator := &activator{
		Config: config,

		ID:   id.MustNew(),
		Type: ObjectType,
	}

	if newActivator.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newActivator.FactoryCollection == nil {
		return nil, maskAnyf(invalidConfigError, "factory collection must not be empty")
	}
	if newActivator.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}

	newActivator.Log.Register(newActivator.GetType())

	return newActivator, nil
}

// MustNew creates either a new default configured activator object, or panics.
func MustNew() spec.Activator {
	newActivator, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newActivator
}

type activator struct {
	Config

	ID   spec.ObjectID
	Type spec.ObjectType
}

func (a *activator) Activate(CLG spec.CLG, networkPayload spec.NetworkPayload) (spec.NetworkPayload, error) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call Activate")

	// Fetch the queued network payloads. queue is a string of comma separated
	// JSON objects representing a specific network payload.
	behaviourID, ok := networkPayload.GetContext().GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	queueKey := key.NewCLGKey("activate:queue:behaviour-id:%s:network-payload", behaviourID)
	s, err := a.Storage().General().Get(queueKey)
	if err != nil {
		return nil, maskAny(err)
	}
	queue, err := stringToQueue(s)
	if err != nil {
		return nil, maskAny(err)
	}

	// Merge the given network payload with the queue that we just fetched from
	// storage. We store the extended queue directly after merging it with the
	// given network payload to definitely track the received network payload,
	// even if something goes wrong and we need to return an error on the code
	// below. In case the current queue exeeds a certain amount of payloads, it is
	// unlikely that the queue is going to be helpful when growing any further.
	// Thus we cut the queue at some point beyond the interface capabilities of
	// the requested CLG. Note that it is possible to have multiple network
	// payloads sent by the same CLG. That might happen in case a specific CLG
	// wants to fulfil the interface of the requested CLG on its own, even it is
	// not able to do so with the output of a single calculation.
	queue = append(queue, networkPayload)
	queueBuffer := len(getInputTypes(CLG.GetCalculate())) + 1
	if len(queue) > queueBuffer {
		queue = queue[1:]
	}
	err = n.persistQueue(queueKey, queue)
	if err != nil {
		return nil, maskAny(err)
	}

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(CLG spec.CLG, queue []spec.NetworkPayload) (spec.NetworkPayload, error){
		a.GetNetworkPayload,
		a.NewNetworkPayload,
	}

	// Execute one lookup after another. As soon as we find a network payload, we
	// return it.
	var newNetworkPayload spec.NewNetworkPayload
	for _, lookup := range lookups {
		newNetworkPayload, err = lookup(CLG, queue)
		if IsNetworkPayloadNotFound(err) {
			// There could no network payload be found by this lookup. Go on and try
			// the next one.
			continue
		} else if err != nil {
			return nil, maskAny(err)
		}

		// The current lookup was successful. We do not need to execute any further
		// lookup, but can go on with the new network payload created.
		break
	}

	// Filter all network payloads from the queue that are merged into the new
	// network payload.
	var newQueue spec.NetworkPayload
	for _, s := range newNetworkPayload.GetSources() {
		for _, np := range queue {
			// At this point there is only one source given. That is the CLG that
			// forwarded the current network payload to here. If this is not the case,
			// we return an error.
			sources := np.GetSources()
			if len(sources) != 1 {
				return nil, maskAnyf(invalidSourcesError, "there must be one source")
			}
			if s == string(sources[0]) {
				// The current network payload is part of the merged network payload.
				// Thus we do not add it to the new queue.
				continue
			}
			newQueue = append(newQueue, np)
		}
	}

	// Update the modified queue in the underlying storage.
	err = n.persistQueue(queueKey, newQueue)
	if err != nil {
		return nil, maskAny(err)
	}

	// The current lookup was able to find a network payload. Thus we simply
	// return it.
	return newNetworkPayload, nil
}

func (a *activator) GetNetworkPayload(CLG spec.CLG, queue []spec.NetworkPayload) (spec.NetworkPayload, error) {
	// Fetch the combination of successful behaviour IDs which are known to be
	// useful for the activation of the requested CLG. The network payloads sent
	// by the CLGs being fetched here are known to be useful because they have
	// already been helpful for the execution of the current CLG tree.
	behaviourID, ok := queue[0].GetContext().GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := key.NewCLGKey("activate:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	s, err := a.Storage().General().Get(behaviourIDsKey)
	if storage.IsNotFound(err) {
		// No successful combination of behaviour IDs is stored. Thus we return an
		// error. Eventually some other lookup is able to find a sufficient network
		// payload.
		return nil, maskAny(networkPayloadNotFoundError)
	} else if err != nil {
		return nil, maskAny(err)
	}
	behaviourIDs := strings.Split(s, ",")
	if len(behaviourIDs) == 0 {
		// No activation configuration of the requested CLG is stored. Thus we
		// return an error. Eventually some other lookup is able to find a
		// sufficient network payload.
		return nil, maskAny(networkPayloadNotFoundError)
	}

	// Check if there is a queued network payload for each behaviour ID we found in the
	// storage. Here it is important to obtain the order of the behaviour IDs
	// stored as connections. They represent the input interface of the requested
	// CLG. Thus there must not be any variation applied to the lookup here,
	// because we need the lookup to be reproducable.
	var matches []spec.NetworkPayload
	for _, behaviourID := range behaviourIDs {
		for _, np := range queue {
			// At this point there is only one source given. That is the CLG that
			// forwarded the current network payload to here. If this is not the case,
			// we return an error.
			sources := np.GetSources()
			if len(sources) != 1 {
				return nil, maskAnyf(invalidSourcesError, "there must be one source")
			}
			if behaviourID == string(sources[0]) {
				// The current behaviour ID belongs to the current network payload. We
				// add the matching network payload to our list and go on to find the
				// network payload belonging to the next behabiour ID.
				matches = append(matches, np)
				break
			}
		}
	}
	if len(behaviourIDs) != len(matches) {
		// No match using the stored configuration associated with the requested CLG
		// can be found. Thus we return an error. Eventually some other lookup is
		// able to find a sufficient network payload.
		return nil, maskAny(networkPayloadNotFoundError)
	}

	// The received network payloads are able to satisfy the interface of the
	// requested CLG. We merge the matching network payloads together and return
	// the result.
	newNetworkPayload, err := mergeNetworkPayloads(matches)
	if err != nil {
		return nil, maskAny(err)
	}

	return newNetworkPayload, nil
}

func (a *activator) NewNetworkPayload(CLG spec.CLG, queue []spec.NetworkPayload) (spec.NetworkPayload, error) {
	// Track the input types of the requested CLG as string slice to have
	// something that is easily comparable and efficient.
	clgTypes := typesToStrings(getInputTypes(CLG.GetCalculate()))

	// Prepare the permutation list to find out which combination of payloads
	// satisfies the requested CLG's interface.
	newPermutationListConfig := permutation.DefaultListConfig()
	newPermutationListConfig.MaxGrowth = len(clgTypes)
	newPermutationListConfig.Values = queueToValues(queue)
	newPermutationList, err := permutation.NewList(newPermutationListConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	// Permute the permutation list of the queued network payloads until we found
	// all the matching combinations.
	var possibleMatches []spec.NetworkPayload
	for {
		// Check if the current combination of network payloads already satisfies
		// the interface of the requested CLG. This is done in the first place to
		// also handle the very first combination of the permutation list.  In case
		// there does a combination of network payloads match the interface of the
		// requested CLG, we capture the found combination and try to find more
		// combinations in the upcoming loops.
		permutedValues := newPermutationList.GetPermutedValues()
		valueTypes := typesToStrings(valuesToTypes(permutedValues))
		if equalStrings(clgTypes, valueTypes) {
			possibleMatches = valuesToQueue(permutedValues)
		}

		// Permute the list of the queued network payloads by one further
		// permutation step within the current iteration. As soon as the permutation
		// list cannot be permuted anymore, we stop the permutation loop to choose
		// one random combination of the tracked list in the next step below.
		err = a.Factory().Permutation().PermuteBy(newPermutationList, 1)
		if IsMaxGrowthReached(err) {
			break
		} else if err != nil {
			return nil, maskAny(err)
		}
	}

	// We fetched all possible combinations if network payloads that match the
	// interface of the requested CLG. Now we need to select one random
	// combination to cover all possible combinations across all possible CLG
	// trees being created over time. This prevents us from choosing always only
	// the first matching combination, which would lack discoveries of all
	// potential combinations being created.
	matchIndex, err := a.Factory().Random().CreateMax(len(possibleMatches))
	if err != nil {
		return nil, maskAny(err)
	}
	matches := possibleMatches[matchIndex]

	// The queued network payloads are able to satisfy the interface of the
	// requested CLG. We merge the matching network payloads together and return
	// the result after storing the created configuration of the requested CLG.
	newNetworkPayload, err := mergeNetworkPayloads(matches)
	if err != nil {
		return nil, maskAny(err)
	}

	// Persists the combination of permuted network payloads as configuration for
	// the requested CLG. This configuration is stored using references of the
	// behaviour IDs associated with CLGs that forwarded signals to this requested
	// CLG. Note that the order of behaviour IDs must be preserved, because it
	// represents the input interface of the requested CLG.
	behaviourID, ok := newNetworkPayload.GetContext().GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := key.NewCLGKey("activate:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	var behaviourIDs []string
	for _, behaviourID := range newNetworkPayload.GetSources() {
		behaviourIDs = append(behaviourIDs, string(behaviourID))
	}
	err := a.Storage().General().Set(behaviourIDsKey, strings.Join(behaviourIDs, ","))
	if err != nil {
		return nil, maskAny(err)
	}

	return newNetworkPayload, nil
}
