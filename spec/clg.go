package spec

// CLG represents the name of a single method implementing basic behavior. CLGs
// implemented in the clg package are supposed to be used as CLG type. Thus
// each CLG can be executed using the clg package's Execute function that takes
// the CLG name as input argument.
type CLG string
