package event

var (
	ExtendTailType Type = "extend-tail"
)

// Config represents the configuration used to create a new queue object.
type ExtendTailConfig struct {
	// Connection represents the new connection being tracked during the current
	// event. This connection consist out of two peers. The first peer is
	// Destination. The second peer is Source.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the event being tracked. In this case,
	// ExtendTailType.
	ConnectionPath string

	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents one source of the network payload currently being
	// processed.
	Source string
}

// DefaultEventQueueConfig provides a default configuration to create a new
// extend tail object by best effort.
func DefaultExtendTailConfig() ExtendTailConfig {
	newConfig := ExtendTailConfig{
		Connection:     "",
		ConnectionPath: "",
		Destination:    "",
		Source:         "",
	}

	return newConfig
}

// NewExtendTail creates a new configured extend tail object.
func NewExtendTail(config ExtendTailConfig) (Event, error) {
	newEvent := &extendTail{
		ExtendTailConfig: config,

		Type: ExtendTailType,
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

type extendTail struct {
	ExtendTailConfig

	Type Type
}

func (et *extendTail) GetConnection() string {
	return et.Connection
}

func (et *extendTail) GetConnectionPath() string {
	return et.ConnectionPath
}

func (et *extendTail) GetDestination() string {
	return et.Destination
}

func (et *extendTail) GetSource() string {
	return et.Source
}

func (et *extendTail) GetType() Type {
	return et.Type
}
