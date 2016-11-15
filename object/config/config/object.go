package config

// New creates a new config object. It provides configuration about the config
// file.
func New() *Object {
	return &Object{}
}

// Object represents the config file config object.
type Object struct {
	// Settings.

	// dir represents the directory in which the config file can be found.
	dir string
	// name represents the file name of the config file without extension. The
	// actual config file can have either json or yaml extension and format.
	name string
}

// Dir returns the dir of the file config.
func (o *Object) Dir() string {
	return o.dir
}

// Name returns the name of the file config.
func (o *Object) Name() string {
	return o.name
}

// SetDir sets the dir for the file config.
func (o *Object) SetDir(dir *string) {
	o.dir = *dir
}

// SetName sets the name for the file config.
func (o *Object) SetName(name *string) {
	o.name = *name
}
