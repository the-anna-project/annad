package activator

import (
	"encoding/json"
	"reflect"
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
	StorageCollection spec.StorageCollection
}

// DefaultConfig provides a default configuration to create a new activator
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log:               log.New(log.DefaultConfig()),
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
	behaviourID, ok := networkPayload.GetContext().GetBehaviorID()
	if !ok {
		return nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}
	queueKey := key.NewCLGKey("activate:queue:behaviour-id:%s:network-payload", behaviourID)
	s, err := a.Storage().General().Get(queueKey)
	if err != nil {
		return nil, maskAny(err)
	}
	var queue []spec.NetworkPayload
	for _, s := range strings.Split(s, ",") {
		np := api.MustNewNetworkPayload()
		err = json.Unmarshal([]byte(s), &np)
		if err != nil {
			return nil, maskAny(err)
		}
		queue = append(queue, np)
	}

	// Merge the given payload with the fetched queue. Note that it is possible to
	// have multiple network payloads sent by the same CLG. That might happen in
	// case a specific CLG wants to fulfil the interface of the requested CLG on
	// its own, even it is not able to do so with the output of a single
	// calculation. We store the extended queue directly after merging to
	// definitely have tracked the received network payload, even if something
	// goes wrong and we need to return an error on the code below.
	queue = append(queue, networkPayload)
	raw, err = json.marshal(queue)
	if err != nil {
		return nil, maskAny(err)
	}
	err := a.Storage().General().Set(queueKey, string(raw))
	if err != nil {
		return nil, maskAny(err)
	}

	// This is the list of lookup functions which is executed seuqentially.
	lookups := []func(networkPayload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error){
		a.GetNetworkPayload,
		a.NewNetworkPayload,
	}

	// Execute one lookup after another. As soon as we find a network payload, we
	// return it.
	for _, lookup := range lookups {
		newNetworkPayload, newQueue, err = lookup(networkPayload, queue)
		if IsNetworkPayloadNotFound(err) {
			// There could no network payload be found by this lookup. Go on and try
			// the next one.
			continue
		} else if err != nil {
			return nil, maskAny(err)
		}

		// Update the modified queue.
		raw, err = json.marshal(newQueue)
		if err != nil {
			return nil, maskAny(err)
		}
		err := a.Storage().General().Set(queueKey, string(raw))
		if err != nil {
			return nil, maskAny(err)
		}

		// The current lookup was able to find a network payload. Thus we simply
		// return it.
		return newNetworkPayload, nil
	}

	return maskAny(networkPayloadNotFoundError), nil
}

func (a *activator) GetNetworkPayload(networkPayload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	// Fetch the combination of successful behaviour IDs which are known to be
	// useful for the activation of the requested CLG. The network payloads sent
	// by the CLGs being fetched here are known to be useful because they have
	// already been helpful for the execution of the current CLG tree.
	behaviourID, ok := networkPayload.GetContext().GetBehaviorID()
	if !ok {
		return nil, nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}
	// TODO this data needs to be written when creating new combinations
	behaviourIDsKey := key.NewCLGKey("activate:success:behaviour-id:%s:behaviour-ids", behaviourID)
	s, err := a.Storage().General().Get(behaviourIDsKey)
	if storage.IsNotFound(err) {
		// No successful combination of behaviour IDs is stored. Thus we return an
		// error. Eventually some other lookup is able to find a sufficient network
		// payload.
		return nil, nil, maskAny(networkPayloadNotFoundError)
	} else if err != nil {
		return nil, nil, maskAny(err)
	}
	behaviourIDs := strings.Split(s, ",")
	if len(behaviourIDs) == 0 {
		// No successful combination of behaviour IDs is stored. Thus we return an
		// error. Eventually some other lookup is able to find a sufficient network
		// payload.
		return nil, nil, maskAny(networkPayloadNotFoundError)
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
				return nil, nil, maskAnyf(invalidSourcesError, "there must be one source")
			}
			if behaviourID == string(sources[0]) {
				// The current behaviour ID belongs to the current network payload. We
				// remove the network payload from the queue and try to find the network
				// payload belonging to the next behabiour ID.
				matches = append(matches, np)
				queue = filterNetworkPayload(queue, np)
				break
			}
		}
	}
	if len(behaviourIDs) != len(matches) {
		// No match using the stored configuration associated with the requested CLG
		// can be found. Thus we return an error. Eventually some other lookup is
		// able to find a sufficient network payload.
		return nil, nil, maskAny(networkPayloadNotFoundError)
	}

	// The received network payloads have been able to satisfy the interface of
	// the requested CLG. We merge the matching network payloads together and
	// return the result.
	newNetworkPayload, err := mergeNetworkPayloads(matches)
	if err != nil {
		return nil, nil, maskAny(err)
	}

	return newNetworkPayload, queue, nil
}

// TODO
func (a *activator) NewNetworkPayload(networkPayload spec.NetworkPayload, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	// fetch queue
	// merge given payload with queue
}

// TODO

// payloadFromPermutations tries to find a combination of payloads which
// together are able to fulfill the interface of the requested CLG. The first
// combination of payloads found, which match the interface of the requested CLG
// will be returned in one merged network payload. The returned list of network
// payloads will not contain any of the network payloads merged.
func (a *activator) payloadFromPermutations(ctx spec.Context, queue []spec.NetworkPayload) (spec.NetworkPayload, []spec.NetworkPayload, error) {
	clgName, ok := ctx.GetCLGName()
	if !ok {
		return nil, nil, maskAnyf(invalidCLGNameError, "must not be empty")
	}
	CLG, err := a.clgByName(clgName)
	if err != nil {
		return nil, nil, maskAny(err)
	}
	inputTypes := CLG.GetInputTypes()

	// Prepare the permutation list to find out which combination of payloads
	// satisfies the requested CLG's interface.
	newConfig := permutation.DefaultListConfig()
	newConfig.MaxGrowth = len(inputTypes)
	newConfig.Values = queueToValues(queue)
	newPermutationList, err := permutation.NewList(newConfig)
	if err != nil {
		return nil, nil, maskAny(err)
	}

	for {
		err := a.Factory().Permutation().MapTo(newPermutationList)
		if err != nil {
			return nil, nil, maskAny(err)
		}

		// Check if the given payload satisfies the requested CLG's interface.
		members := newPermutationList.GetMembers()
		types, err := membersToTypes(members)
		if err != nil {
			return nil, nil, maskAny(err)
		}
		if reflect.DeepEqual(types, inputTypes) {
			newPayload, err := membersToPayload(ctx, members)
			if err != nil {
				return nil, nil, maskAny(err)
			}
			newQueue, err := filterMembersFromQueue(members, queue)
			if err != nil {
				return nil, nil, maskAny(err)
			}

			return newPayload, newQueue, nil
		}

		err = a.Factory().Permutation().PermuteBy(newPermutationList, 1)
		if err != nil {
			// Note that also an error is thrown when the maximum growth of the
			// permutation list was reached.
			return nil, nil, maskAny(err)
		}
	}
}

// helper

func filterNetworkPayload(list []spec.NetworkPayload, item spec.NetworkPayload) []spec.NetworkPayload {
	var newList spec.NetworkPayload

	for _, np := range list {
		if np.GetID() == item.GetID() {
			// This is the network payload we want to filter. Thus we go on with the
			// loop to not add it to the new list.
			continue
		}

		newList = append(newList, np)
	}

	return newList
}

func containsNetworkPayload(list []spec.NetworkPayload, item spec.NetworkPayload) bool {
	for _, p := range list {
		if p.GetID() == item.GetID() {
			return true
		}
	}

	return false
}

func filterMembersFromQueue(members []interface{}, queue []spec.NetworkPayload) ([]spec.NetworkPayload, error) {
	var memberPayloads []spec.NetworkPayload
	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}
		memberPayloads = append(memberPayloads, payload)
	}

	var newQueue []spec.NetworkPayload
	for _, queuedPayload := range queue {
		if containsNetworkPayload(memberPayloads, queuedPayload) {
			continue
		}

		newQueue = append(newQueue, queuedPayload)
	}

	return newQueue, nil
}

func getInputTypes(f interface{}) []reflect.Type {
	t := reflect.TypeOf(f)

	var inputType []reflect.Type

	for i := 0; i < t.NumIn(); i++ {
		inputType = append(inputType, t.In(i))
	}

	return inputType
}

func mergeNetworkPayloads(networkPayloads []spec.NetworkPayload) (spec.NetworkPayload, error) {
	if len(networkPayloads) == 0 {
		return nil, maskAny(networkPayloadNotFoundError)
	}

	var args []reflect.Value
	var sources []spec.ObjectID
	var ctx = networkPayloads[0].GetContext()

	behaviourID, ok := ctx.GetBehaviorID()
	if !ok {
		return nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}

	for _, m := range members {
		for _, v := range payload.GetArgs() {
			args = append(args, v)
		}

		sources = append(sources, payload.GetSources()...)
	}

	networkPayloadConfig := api.DefaultNetworkPayloadConfig()
	networkPayloadConfig.Args = args
	networkPayloadConfig.Context = ctx
	networkPayloadConfig.Destination = spec.ObjectID(behaviourID)
	networkPayloadConfig.Sources = sources
	networkPayload, err := api.NewNetworkPayload(networkPayloadConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return networkPayload, nil
}

func membersToPayload(ctx spec.Context, members []interface{}) (spec.NetworkPayload, error) {
	var args []reflect.Value
	var sources []spec.ObjectID

	behaviourID, ok := ctx.GetBehaviorID()
	if !ok {
		return nil, maskAnyf(invalidBehaviorIDError, "must not be empty")
	}

	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}

		for _, v := range payload.GetArgs() {
			args = append(args, v)
		}

		sources = append(sources, payload.GetSources()...)
	}

	newPayloadConfig := api.DefaultNetworkPayloadConfig()
	newPayloadConfig.Args = args
	newPayloadConfig.Context = ctx
	newPayloadConfig.Destination = spec.ObjectID(behaviourID)
	newPayloadConfig.Sources = sources
	newPayload, err := api.NewNetworkPayload(newPayloadConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newPayload, nil
}

func membersToTypes(members []interface{}) ([]reflect.Type, error) {
	var types []reflect.Type

	for _, m := range members {
		payload, ok := m.(spec.NetworkPayload)
		if !ok {
			return nil, maskAnyf(invalidInterfaceError, "member must be spec.NetworkPayload")
		}

		for _, v := range payload.GetArgs() {
			types = append(types, v.Type())
		}
	}

	return types, nil
}

func queueToValues(queue []spec.NetworkPayload) []interface{} {
	var values []interface{}

	for _, p := range queue {
		values = append(values, p)
	}

	return values
}
