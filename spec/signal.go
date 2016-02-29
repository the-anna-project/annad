package spec

// Signal represents a signal that can be send and received using a gateway to
// exchange information in a decoupled way.
type Signal interface {
	// GetError returns the error set to the signal. Setting an error to a signal
	// on the receiver side, lets the sending side know that something went
	// wrong.
	GetError() error

	// GetID returns the signal's ID.
	GetID() string

	// GetInput returns the signal's input.
	GetInput() interface{}

	// GetOutput returns the signal's output.
	GetOutput() interface{}

	// GetResponder returns the signal's responder. The reciving side uses the
	// responder to send the signal back to its origin.
	GetResponder() chan Signal

	// SetError associates an error to the signal. See GetError.
	SetError(err error)

	// SetID sets the signal's ID. This can be used to make the signal
	// identifiable. When thinking about job execution the signal's ID can be set
	// to the job ID so upcoming layers know what this signal is about.
	//
	// TODO should there simply be a job ID getter/setter?
	SetID(ID string)

	// SetInput sets the signal's input.
	SetInput(input interface{})

	// SetOutput sets the signal's output.
	SetOutput(output interface{})
}
