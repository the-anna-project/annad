package session

// New creates a new session object. It provides configuration about the session
// file.
func New() *Object {
	return &Object{}
}

// Object represents the session file config object.
type Object struct {
	// Settings.

	// id represents the ID of the session.
	id string
}

// ID returns the id of the session.
func (o *Object) ID() string {
	return o.id
}

// SetID sets the id for the file config.
func (o *Object) SetID(id *string) {
	o.id = *id
}
