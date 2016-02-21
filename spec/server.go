package spec

type Server interface {
	// Boot initializes and starts the whole server like booting a machine. The
	// call to Boot blocks forever.
	Boot()

	Object
}
