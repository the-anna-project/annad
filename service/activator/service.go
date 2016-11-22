package activator

import (
	"fmt"
	"strings"

	permutationlist "github.com/the-anna-project/permutation/object/list"
	"github.com/the-anna-project/permutation/service"
	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/the-anna-project/annad/service/storage"
)

// New creates a new activator service.
func New() servicespec.ActivatorService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "activator",
		"type": "service",
	}
}

func (s *service) Activate(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
	s.Service().Log().Line("func", "Activate")

	// Fetch the queued network payloads. queue is s string of comma separated
	// JSON objects representing s specific network payload.
	behaviourID, ok := networkPayload.GetContext().GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	queueKey := fmt.Sprintf("activate:queue:behaviour-id:%s:network-payload", behaviourID)
	str, err := s.Service().Storage().General().Get(queueKey)
	if err != nil {
		return nil, maskAny(err)
	}
	queue, err := stringToQueue(str)
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
	err = s.persistQueue(queueKey, queue)
	if err != nil {
		return nil, maskAny(err)
	}

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(CLG servicespec.CLGService, queue []objectspec.NetworkPayload) (objectspec.NetworkPayload, error){
		s.GetNetworkPayload,
		s.New,
	}

	// Execute one lookup after another. As soon as we find a network payload, we
	// return it.
	var newNetworkPayload objectspec.NetworkPayload
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
		// lookup, but can go on with the network payload found.
		break
	}

	// Filter all network payloads from the queue that are merged into the new
	// network payload.
	var newQueue []objectspec.NetworkPayload
	for _, s := range newNetworkPayload.GetSources() {
		for _, np := range queue {
			// At this point there is only one source given. That is the CLG that
			// forwarded the current network payload to here. If this is not the case,
			// we return an error.
			sources := np.GetSources()
			if len(sources) != 1 {
				return nil, maskAnyf(invalidSourcesError, "there must be one source")
			}
			if s == sources[0] {
				// The current network payload is part of the merged network payload.
				// Thus we do not add it to the new queue.
				continue
			}
			newQueue = append(newQueue, np)
		}
	}

	// Update the modified queue in the underlying storage.
	err = s.persistQueue(queueKey, newQueue)
	if err != nil {
		return nil, maskAny(err)
	}

	// The current lookup was able to find a network payload. Thus we simply
	// return it.
	return newNetworkPayload, nil
}

func (s *service) GetNetworkPayload(CLG servicespec.CLGService, queue []objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
	// Fetch the combination of successful behaviour IDs which are known to be
	// useful for the activation of the requested CLG. The network payloads sent
	// by the CLGs being fetched here are known to be useful because they have
	// already been helpful for the execution of the current CLG tree.
	behaviourID, ok := queue[0].GetContext().GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := fmt.Sprintf("activate:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	str, err := s.Service().Storage().General().Get(behaviourIDsKey)
	if storage.IsNotFound(err) {
		// No successful combination of behaviour IDs is stored. Thus we return an
		// error. Eventually some other lookup is able to find a sufficient network
		// payload.
		return nil, maskAny(networkPayloadNotFoundError)
	} else if err != nil {
		return nil, maskAny(err)
	}
	behaviourIDs := strings.Split(str, ",")
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
	// because we need the lookup to be reproducible.
	var matches []objectspec.NetworkPayload
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

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) New(CLG servicespec.CLGService, queue []objectspec.NetworkPayload) (objectspec.NetworkPayload, error) {
	// Track the input types of the requested CLG as string slice to have
	// something that is easily comparable and efficient. By convention the first
	// input argument of each CLG is a context. We remove the first argument here,
	// because we want to match output interfaces against input interfaces. By
	// convention output interfaces of CLGs must not have a context as first
	// return value. Therefore we align the input and output values to make them
	// comparable.
	clgTypes := typesToStrings(getInputTypes(CLG.GetCalculate()))[1:]

	// Prepare the permutation list to find out which combination of payloads
	// satisfies the requested CLG's interface.
	permutationList := permutationlist.New()
	permutationList.SetMaxGrowth(len(clgTypes))
	permutationList.SetRawValues(queueToValues(queue))

	// Permute the permutation list of the queued network payloads until we found
	// all the matching combinations.
	var possibleMatches [][]objectspec.NetworkPayload
	for {
		// Check if the current combination of network payloads already satisfies
		// the interface of the requested CLG. This is done in the first place to
		// also handle the very first combination of the permutation list.  In case
		// there does a combination of network payloads match the interface of the
		// requested CLG, we capture the found combination and try to find more
		// combinations in the upcoming loops.
		permutedValues := permutationList.PermutedValues()
		valueTypes := typesToStrings(valuesToTypes(permutedValues))
		if equalStrings(clgTypes, valueTypes) {
			possibleMatches = append(possibleMatches, valuesToQueue(permutedValues))
		}

		// Permute the list of the queued network payloads by one further
		// permutation step within the current iteration. As soon as the permutation
		// list cannot be permuted anymore, we stop the permutation loop to choose
		// one random combination of the tracked list in the next step below.
		err := s.Service().Permutation().PermuteBy(permutationList, 1)
		if permutation.IsMaxGrowthReached(err) {
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
	matchIndex, err := s.Service().Random().CreateMax(len(possibleMatches))
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
	behaviourIDsKey := fmt.Sprintf("activate:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	var behaviourIDs []string
	for _, behaviourID := range newNetworkPayload.GetSources() {
		behaviourIDs = append(behaviourIDs, string(behaviourID))
	}
	err = s.Service().Storage().General().Set(behaviourIDsKey, strings.Join(behaviourIDs, ","))
	if err != nil {
		return nil, maskAny(err)
	}

	return newNetworkPayload, nil
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}
