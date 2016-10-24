package event

var (
	NewPathType Type = "new-path"
)

// Config represents the configuration used to create a new queue object.
type NewPathConfig struct {
	// Connection represents the new connection being tracked during the current
	// event. This connection consist out of two peers. The first peer is
	// Destination. The second peer is Source.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the event being tracked. In this case,
	// NewPathType.
	ConnectionPath string

	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents one source of the network payload currently being
	// processed.
	Source string
}

// DefaultEventQueueConfig provides a default configuration to create a new
// new path object by best effort.
func DefaultNewPathConfig() NewPathConfig {
	newConfig := NewPathConfig{
		Connection:     "",
		ConnectionPath: "",
		Destination:    "",
		Source:         "",
	}

	return newConfig
}

// NewNewPath a new configured new path object.
func NewNewPath(config NewPathConfig) (Event, error) {
	newEvent := &newPath{
		NewPathConfig: config,

		Type: NewPathType,
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

type newPath struct {
	NewPathConfig

	Type Type
}

func (mp *newPath) GetConnection() string {
	return mp.Connection
}

func (mp *newPath) GetConnectionPath() string {
	return mp.ConnectionPath
}

func (mp *newPath) GetDestination() string {
	return mp.Destination
}

func (mp *newPath) GetSource() string {
	return mp.Source
}

func (mp *newPath) GetType() Type {
	return mp.Type
}
