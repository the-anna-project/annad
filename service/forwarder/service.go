package forwarder

import (
	"encoding/json"
	"fmt"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/the-anna-project/annad/object/networkpayload"
	"github.com/the-anna-project/annad/service/storage"
)

// New creates a new forwarder service.
func New() servicespec.ForwarderService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata map[string]string
	// maxSignals represents the maximum number of signals being forwarded by one
	// CLG. When a requested CLG needs to decide where to forward signals to, it
	// may will forward up to maxSignals signals to other CLGs, if any.
	maxSignals int
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"name": "forwarder",
		"type": "service",
	}

	s.maxSignals = 5
}

func (s *service) Forward(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) error {
	s.Service().Log().Line("func", "Forward")

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) ([]objectspec.NetworkPayload, error){
		s.GetNetworkPayloads,
		s.News,
	}

	// Execute one lookup after another. As soon as we find some behaviour IDs, we
	// use them to forward the given network payload.
	var newNetworkPayloads []objectspec.NetworkPayload
	var err error
	for _, lookup := range lookups {
		newNetworkPayloads, err = lookup(CLG, networkPayload)
		if IsNetworkPayloadsNotFound(err) {
			// There could no behaviour IDs be found by this lookup. Go on and try the
			// next one.
			continue
		} else if err != nil {
			return maskAny(err)
		}

		// The current lookup was successful. We do not need to execute any further
		// lookup, but can go on with the behaviour IDs found.
		break
	}

	// Forward the found network payloads to other CLGs by adding them to the
	// queue so other processes can fetch them.
	for _, np := range newNetworkPayloads {
		networkPayloadKey := fmt.Sprintf("events:network-payload")
		b, err := json.Marshal(np)
		if err != nil {
			return maskAny(err)
		}
		// TODO store asynchronuously
		err = s.Service().Storage().General().PushToSet(networkPayloadKey, string(b))
		if err != nil {
			return maskAny(err)
		}
	}

	return nil
}

func (s *service) GetMaxSignals() int {
	return s.maxSignals
}

func (s *service) GetNetworkPayloads(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) ([]objectspec.NetworkPayload, error) {
	ctx := networkPayload.GetContext()

	// Check if there are behaviour IDs known that we can use to forward the
	// current network payload to.
	behaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := fmt.Sprintf("forward:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	newBehaviourIDs, err := s.Service().Storage().General().GetAllFromSet(behaviourIDsKey)
	if storage.IsNotFound(err) {
		// No configuration of behaviour IDs is stored. Thus we return an error.
		// Eventually some other lookup is able to find sufficient network payloads.
		return nil, maskAny(networkPayloadsNotFoundError)
	} else if err != nil {
		return nil, maskAny(err)
	}

	// Create a list of new network payloads.
	var newNetworkPayloads []objectspec.NetworkPayload
	for _, behaviourID := range newBehaviourIDs {
		// Prepare a new context for the new network payload.
		newCtx := ctx.Clone()
		newCtx.SetBehaviourID(behaviourID)

		// Create a new network payload.
		newNetworkPayloadConfig := networkpayload.DefaultConfig()
		newNetworkPayloadConfig.Args = networkPayload.GetArgs()
		newNetworkPayloadConfig.Context = newCtx
		newNetworkPayloadConfig.Destination = string(behaviourID)
		newNetworkPayloadConfig.Sources = []string{networkPayload.GetDestination()}
		newNetworkPayload, err := networkpayload.New(newNetworkPayloadConfig)
		if err != nil {
			return nil, maskAny(err)
		}

		newNetworkPayloads = append(newNetworkPayloads, newNetworkPayload)
	}

	return newNetworkPayloads, nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) News(CLG servicespec.CLGService, networkPayload objectspec.NetworkPayload) ([]objectspec.NetworkPayload, error) {
	ctx := networkPayload.GetContext()

	// Decide how many new behaviour IDs should be created. This defines the
	// number of signals being forwarded to other CLGs. Here we want to make a
	// pseudo random decision. CreateMax takes a max paramater which is exclusive.
	// Therefore we increment the configuration for the maximum signals desired by
	// one, to reflect the maximum setting properly.
	maxSignals, err := s.Service().Random().CreateMax(s.GetMaxSignals() + 1)
	if err != nil {
		return nil, maskAny(err)
	}

	// Create the desired number of behaviour IDs.
	var newBehaviourIDs []string
	for i := 0; i < maxSignals; i++ {
		newBehaviourID, err := s.Service().ID().New()
		if err != nil {
			return nil, maskAny(err)
		}
		newBehaviourIDs = append(newBehaviourIDs, string(newBehaviourID))
	}

	// TODO find a CLG name that can be connected to the current CLG for each new
	// behaviour ID and pair these combinations (network event tracker)

	// Store each new behaviour ID in the underlying storage.
	behaviourID, ok := ctx.GetBehaviourID()
	if !ok {
		return nil, maskAnyf(invalidBehaviourIDError, "must not be empty")
	}
	behaviourIDsKey := fmt.Sprintf("forward:configuration:behaviour-id:%s:behaviour-ids", behaviourID)
	for _, behaviourID := range newBehaviourIDs {
		// TODO store asynchronuously
		err = s.Service().Storage().General().PushToSet(behaviourIDsKey, behaviourID)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	// Create a list of new network payloads.
	var newNetworkPayloads []objectspec.NetworkPayload
	for _, behaviourID := range newBehaviourIDs {
		// Prepare a new context for the new network payload.
		newCtx := ctx.Clone()
		newCtx.SetBehaviourID(behaviourID)
		// TODO set the paired CLG name to the new context

		// Create a new network payload.
		newNetworkPayloadConfig := networkpayload.DefaultConfig()
		newNetworkPayloadConfig.Args = networkPayload.GetArgs()
		newNetworkPayloadConfig.Context = newCtx
		newNetworkPayloadConfig.Destination = string(behaviourID)
		newNetworkPayloadConfig.Sources = []string{networkPayload.GetDestination()}
		newNetworkPayload, err := networkpayload.New(newNetworkPayloadConfig)
		if err != nil {
			return nil, maskAny(err)
		}

		newNetworkPayloads = append(newNetworkPayloads, newNetworkPayload)
	}

	return newNetworkPayloads, nil
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}
