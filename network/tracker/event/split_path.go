package event

var (
	SplitPathType Type = "split-path"
)

// SplitPathConfig represents the configuration used to create a new queue
// object.
type SplitPathConfig struct {
	// Connection represents the new connection being tracked during the current
	// event. This connection consist out of two peers. The first peer is
	// Destination. The second peer is Source.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the event being tracked. In this case,
	// SplitPathType.
	ConnectionPath string

	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents one source of the network payload currently being
	// processed.
	Source string
}

// DefaultEventQueueConfig provides a default configuration to create a new
// split path object by best effort.
func DefaultSplitPathConfig() SplitPathConfig {
	newConfig := SplitPathConfig{
		Connection:     "",
		ConnectionPath: "",
		Destination:    "",
		Source:         "",
	}

	return newConfig
}

// NewSplitPath a new configured split path object.
func NewSplitPath(config SplitPathConfig) (Event, error) {
	newEvent := &splitPath{
		SplitPathConfig: config,

		Type: SplitPathType,
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

type splitPath struct {
	SplitPathConfig

	Type Type
}

func (mp *splitPath) GetConnection() string {
	return mp.Connection
}

func (mp *splitPath) GetConnectionPath() string {
	return mp.ConnectionPath
}

func (mp *splitPath) GetDestination() string {
	return mp.Destination
}

func (mp *splitPath) GetSource() string {
	return mp.Source
}

func (mp *splitPath) GetType() Type {
	return mp.Type
}
