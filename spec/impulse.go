package spec

// Impulse is basically a container to carry around information. It walks
// through neural networks and is modified on its way to finally creating some
// output based on some input.
//
// An impulse carries information about former impulses and their corresponding
// inputs around. This enables the network to collect related information to
// make a reasonable point about the current contextual task.
//
//
type Impulse interface {
	// GetActions returns the currently configured action items that are used to
	// create a reasonable strategy.
	GetActions() []ObjectType

	// GetAllInputs returns the key-value mapping of impulse IDs and their
	// corresponding inputs for the current impulse, if any. Here we are able to
	// collect inputs related to the session. The structure looks as follows. key
	// is the session ID. value is the corresponding input.
	//
	//     key: value
	//     key: value
	//     ...
	//
	GetAllInputs() map[ObjectID]string

	// GetAllStrategies returns the key-value mapping of requestor object types
	// and their corresponding generated stratgies the current impulse has
	// tracked to far, if any. Here we are able to collect strategies related to
	// the walk of the current impulse. The structure looks as follows. key is
	// the requestor's object type. value is the corresponding strategy.
	//
	//     key: value
	//     key: value
	//     ...
	//
	GetAllStrategies() map[ObjectType]Strategy

	// GetInputByImpulseID returns the input known to the impulse under the given
	// impulse ID, if any. If there cannot be found any input for the given
	// impulse ID, an error, indicating that, is returned.
	GetInputByImpulseID(impulseID ObjectID) (string, error)

	// GetOutput returns the impulse's output.
	GetOutput() string

	// GetRequestor returns the currently configured requestor.
	GetRequestor() ObjectType

	// GetSessionID returns the ID of the session related to the current impulse.
	GetSessionID() string

	// GetStrategy returns the currently configured strategy of the impulse.
	GetStrategy() Strategy

	Object

	// SetActions stores the given actions to the impulse. This actions are used
	// to create reasonable strategies. When setting new actions here, the
	// formerly set actions will be overwritten.
	SetActions(actions []ObjectType)

	// SetInputByImpulseID causes the impulse to store the given input using the
	// given impulse ID in memory.
	SetInputByImpulseID(impulseID ObjectID, input string)

	// SetRequestor stores the given requestor to the impulse.
	SetRequestor(requestor ObjectType)

	// SetOutput sets the impulse's output.
	SetOutput(output string)

	// SetStrategy stores the given strategy to the impulse using the given
	// requestor.
	SetStrategyByRequestor(requestor ObjectType, strategy Strategy)
}
