// Package service implements a service to manage connections inside network
// layers.
package service

import (
	"fmt"
	"strings"

	peerservice "github.com/the-anna-project/peer/service"
	servicespec "github.com/the-anna-project/spec/service"
)

const (
	KindBehaviour   = "behaviour"
	KindInformation = "information"
	KindPosition    = "position"
)

// Config represents the configuration used to create a new layer service.
type Config struct {
	// Settings.
	Kind string
}

// DefaultConfig provides a default configuration to create a new layer service
// by best effort.
func DefaultConfig() Config {
	return Config{
		// Settings.
		Kind: "",
	}
}

// New creates a new layer service.
func New(config Config) (servicespec.LayerService, error) {
	// Settings.
	if config.Kind == "" {
		return nil, maskAnyf(invalidConfigError, "kind must not be empty")
	}

	newService := &service{
		// Dependencies.
		serviceCollection: nil,

		// Internals.
		closer:   make(chan struct{}, 1),
		metadata: map[string]string{},

		// Settings.
		kind: config.Kind,
	}

	return newService, nil
}

type service struct {
	// Dependencies.
	serviceCollection servicespec.ServiceCollection

	// Internals.
	// TODO add Shutdown
	closer   chan struct{}
	metadata map[string]string

	// Settings.
	kind string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"kind": s.Kind(),
		"name": "layer",
		"type": "service",
	}
}

func (s *service) CreatePeer(peer string) (string, error) {
	s.Service().Log().Line("func", "CreatePeer")

	// Define a proper peer key for the peer that is requested to be created.
	peerKey := s.PeerKey(peer)

	// Define the peer creation action.
	createPeer := func(canceler <-chan struct{}) error {
		err := s.Service().Peer().Create(peerKey)
		if peerservice.IsAlreadyExists(err) && s.Kind() == KindPosition {
			// It can happen that peers for position we just came up with already
			// exist. In this case we are already fine and can savely ignore the
			// error to go ahead with creating a connection.
		} else if err != nil {
			return maskAny(err)
		}

		return nil
	}

	// Define the position creation action. Note that a peer's position is a peer
	// itself. Note that we also create the connection between peer and position
	// in here.
	createPosition := func(canceler <-chan struct{}) error {
		position, err := s.Service().Position().Default()
		if err != nil {
			return maskAny(err)
		}

		// Note that CreatePeer creates its own peer key and we must not provide the
		// key here.
		positionKey, err := s.Service().Layer().Position().CreatePeer(position)
		if peerservice.IsAlreadyExists(err) {
			// It can happen that peers for position we just came up with already
			// exist. In this case we are already fine and can savely ignore the
			// error to go ahead with creating a connection.
		} else if err != nil {
			return maskAny(err)
		}

		err = s.Service().Connection().Create(peerKey, positionKey)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	// Define the actions that are going to be executed asynchronously by the
	// worker. In case the executing layer is the position layer itself, we do not
	// want to create a position for the position peer, because a position peer
	// must not have a position peer. Therefore we exclude the action for the
	// position creation from the actions list.
	actions := []func(canceler <-chan struct{}) error{
		createPeer,
	}
	if s.Kind() != KindPosition {
		actions = append(actions, createPosition)
	}

	// Execute the list of actions asynchronously.
	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(actions))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return "", maskAny(err)
	}

	return peerKey, nil
}

func (s *service) DeletePeer(peer string) (string, error) {
	s.Service().Log().Line("func", "DeletePeer")

	// Define a proper peer key for the peer that is requested to be deleted.
	peerKey := s.PeerKey(peer)

	// Define the peer deletion action.
	deletePeer := func(canceler <-chan struct{}) error {
		err := s.Service().Peer().Delete(peerKey)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	// Define the position deletion action. Note that a peer's position is a peer
	// itself. Note that we also delete the connection between peer and position
	// in here.
	deletePosition := func(canceler <-chan struct{}) error {
		position, err := s.Service().Layer().Position().PeerPosition(peerKey)
		if err != nil {
			return maskAny(err)
		}
		// Note that DeletePeer creates its own peer key and we must not provide
		// the key here.
		positionKey, err := s.Service().Layer().Position().DeletePeer(position)
		if err != nil {
			return maskAny(err)
		}
		err = s.Service().Connection().Delete(peerKey, positionKey)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	// Define the actions that are going to be executed asynchronously by the
	// worker. In case the executing layer is the position layer itself, we do not
	// want to delete a position for the position peer, because there is none.
	// Therefore we exclude the action for the position creation from the actions
	// list.
	actions := []func(canceler <-chan struct{}) error{
		deletePeer,
	}
	if s.Kind() != KindPosition {
		actions = append(actions, deletePosition)
	}

	executeConfig := s.Service().Worker().ExecuteConfig()
	executeConfig.SetActions(actions)
	executeConfig.SetCanceler(s.closer)
	executeConfig.SetNumWorkers(len(actions))
	err := s.Service().Worker().Execute(executeConfig)
	if err != nil {
		return "", maskAny(err)
	}

	return peerKey, nil
}

func (s *service) Kind() string {
	return s.kind
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

// PeerKey creates a peer key. Therefore PeerKey must be given a plain peer
// value. The key structure created here is important for layer type
// identification. In case the structure changes here PeerKind below must be
// aligned accordingly to keep the layer implementation functionally correct.
func (s *service) PeerKey(peer string) string {
	return fmt.Sprintf("layer:%s:%s", s.Kind(), peer)
}

// PeerKind tries to return the plain peer kind by parsing the given peer key.
// Therefore it expects a proper key structure as being created by PeerKey. In
// case the peer key structure does not meet the required expectations, PeerKind
// returns an error.
func (s *service) PeerKind(peerKey string) (string, error) {
	splitted := strings.Split(peerKey, ":")

	if len(splitted) != 3 {
		return "", maskAnyf(invalidPeerError, peerKey)
	}
	if splitted[0] != "layer" {
		return "", maskAnyf(invalidPeerError, peerKey)
	}

	return splitted[1], nil
}

func (s *service) PeerPosition(peerKey string) (string, error) {
	peers, err := s.Service().Connection().SearchPeers(peerKey)
	if err != nil {
		return "", maskAny(err)
	}

	var positionKey string
	for _, p := range peers {
		pk, err := s.PeerKind(p)
		if err != nil {
			return "", maskAny(err)
		}
		if pk == KindPosition {
			positionKey = p
			break
		}
	}

	position, err := s.PeerValue(positionKey)
	if err != nil {
		return "", maskAny(err)
	}

	return position, nil
}

// PeerValue tries to return the plain peer value by parsing the given peer key.
// Therefore it expects a proper key structure as being created by PeerKey. In
// case the peer key structure does not meet the required expectations,
// PeerValue returns an error.
func (s *service) PeerValue(peerKey string) (string, error) {
	splitted := strings.Split(peerKey, ":")

	if len(splitted) != 3 {
		return "", maskAnyf(invalidPeerError, peerKey)
	}
	if splitted[0] != "layer" {
		return "", maskAnyf(invalidPeerError, peerKey)
	}

	return splitted[2], nil
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(serviceCollection servicespec.ServiceCollection) {
	s.serviceCollection = serviceCollection
}
