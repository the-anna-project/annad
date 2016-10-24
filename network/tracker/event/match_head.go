package event

var (
	MatchHeadType Type = "match-head"
)

// Config represents the configuration used to create a new queue object.
type MatchHeadConfig struct {
	// Connection represents the new connection being tracked during the current
	// event. This connection consist out of two peers. The first peer is
	// Destination. The second peer is Source.
	Connection string

	// ConnectionPath represents the stored connection path matching the new
	// connection according to the event being tracked. In this case,
	// MatchHeadType.
	ConnectionPath string

	// Destination represents the destination of the network payload currently
	// being processed.
	Destination string

	// Source represents one source of the network payload currently being
	// processed.
	Source string
}

// DefaultEventQueueConfig provides a default configuration to create a new
// match head object by best effort.
func DefaultMatchHeadConfig() MatchHeadConfig {
	newConfig := MatchHeadConfig{
		Connection:     "",
		ConnectionPath: "",
		Destination:    "",
		Source:         "",
	}

	return newConfig
}

// NewMatchHead creates a new configured match head object.
func NewMatchHead(config MatchHeadConfig) (Event, error) {
	newEvent := &matchHead{
		MatchHeadConfig: config,

		Type: MatchHeadType,
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

type matchHead struct {
	MatchHeadConfig

	Type Type
}

func (mh *matchHead) GetConnection() string {
	return mh.Connection
}

func (mh *matchHead) GetConnectionPath() string {
	return mh.ConnectionPath
}

func (mh *matchHead) GetDestination() string {
	return mh.Destination
}

func (mh *matchHead) GetSource() string {
	return mh.Source
}

func (mh *matchHead) GetType() Type {
	return mh.Type
}
