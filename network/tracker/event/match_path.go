package event

var (
	MatchPathType Type = "match-path"
)

// Config represents the configuration used to create a new queue object.
type MatchPathConfig struct {
	// Connection represents the new connection being tracked during the current
	// event. This connection consist out of two peers. The first peer is
	// Destination. The second peer is Source.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the event being tracked. In this case,
	// MatchPathType.
	ConnectionPath string

	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents one source of the network payload currently being
	// processed.
	Source string
}

// DefaultEventQueueConfig provides a default configuration to create a new
// match path object by best effort.
func DefaultMatchPathConfig() MatchPathConfig {
	newConfig := MatchPathConfig{
		Connection:     "",
		ConnectionPath: "",
		Destination:    "",
		Source:         "",
	}

	return newConfig
}

// NewMatchPath creates a new configured match path object.
func NewMatchPath(config MatchPathConfig) (Event, error) {
	newEvent := &matchPath{
		MatchPathConfig: config,

		Type: MatchPathType,
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

type matchPath struct {
	MatchPathConfig

	Type Type
}

func (mp *matchPath) GetConnection() string {
	return mp.Connection
}

func (mp *matchPath) GetConnectionPath() string {
	return mp.ConnectionPath
}

func (mp *matchPath) GetDestination() string {
	return mp.Destination
}

func (mp *matchPath) GetSource() string {
	return mp.Source
}

func (mp *matchPath) GetType() Type {
	return mp.Type
}
