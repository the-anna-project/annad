package event

var (
	MatchTailType Type = "match-tail"
)

// Config represents the configuration used to create a new queue object.
type MatchTailConfig struct {
	// Connection represents the new connection being tracked during the current
	// event. This connection consist out of two peers. The first peer is
	// Destination. The second peer is Source.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the event being tracked. In this case,
	// MatchTailType.
	ConnectionPath string

	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents one source of the network payload currently being
	// processed.
	Source string
}

// DefaultEventQueueConfig provides a default configuration to create a new
// match tail object by best effort.
func DefaultMatchTailConfig() MatchTailConfig {
	newConfig := MatchTailConfig{
		Connection:     "",
		ConnectionPath: "",
		Destination:    "",
		Source:         "",
	}

	return newConfig
}

// NewMatchTail creates a new configured match tail object.
func NewMatchTail(config MatchTailConfig) (Event, error) {
	newEvent := &matchTail{
		MatchTailConfig: config,

		Type: MatchTailType,
	}

	if newEvent.Connection == "" {
		return nil, maskAnyf(invalidConfigError, "connection must not be empty")
	}
	if newEvent.ConnectionPath == "" {
		return nil, maskAnyf(invalidConfigError, "connection path must not be empty")
	}
	if newEvent.Destination == "" {
		return nil, maskAnyf(invalidConfigError, "destination must not be empty")
	}
	if newEvent.Source == "" {
		return nil, maskAnyf(invalidConfigError, "source must not be empty")
	}

	return newEvent, nil
}

type matchTail struct {
	MatchTailConfig

	Type Type
}

func (mp *matchTail) GetConnection() string {
	return mp.Connection
}

func (mp *matchTail) GetConnectionPath() string {
	return mp.ConnectionPath
}

func (mp *matchTail) GetDestination() string {
	return mp.Destination
}

func (mp *matchTail) GetSource() string {
	return mp.Source
}

func (mp *matchTail) GetType() Type {
	return mp.Type
}
