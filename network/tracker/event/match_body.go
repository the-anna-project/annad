package event

var (
	MatchBodyType Type = "match-body"
)

// Config represents the configuration used to create a new queue object.
type MatchBodyConfig struct {
	// Connection represents the new connection being tracked during the current
	// event. This connection consist out of two peers. The first peer is
	// Destination. The second peer is Source.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the event being tracked. In this case,
	// MatchBodyType.
	ConnectionPath string

	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents one source of the network payload currently being
	// processed.
	Source string
}

// DefaultEventQueueConfig provides a default configuration to create a new
// match body object by best effort.
func DefaultMatchBodyConfig() MatchBodyConfig {
	newConfig := MatchBodyConfig{
		Connection:     "",
		ConnectionPath: "",
		Destination:    "",
		Source:         "",
	}

	return newConfig
}

// NewMatchBody creates a new configured match body object.
func NewMatchBody(config MatchBodyConfig) (Event, error) {
	newEvent := &matchBody{
		MatchBodyConfig: config,

		Type: MatchBodyType,
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

type matchBody struct {
	MatchBodyConfig

	Type Type
}

func (mb *matchBody) GetConnection() string {
	return mb.Connection
}

func (mb *matchBody) GetConnectionPath() string {
	return mb.ConnectionPath
}

func (mb *matchBody) GetDestination() string {
	return mb.Destination
}

func (mb *matchBody) GetSource() string {
	return mb.Source
}

func (mb *matchBody) GetType() Type {
	return mb.Type
}
