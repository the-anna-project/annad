package permutation

// TODO
//var (
//	// argTypes represents a list of well known types used to identify CLG input-
//	// and output types. Here we want to have a list of types only.
//	//
//	// TODO identify types by strings
//	argTypes = []interface{}{
//		// Simple types.
//		*new(bool),
//		*new(int),
//		*new(float64),
//		*new(string),
//		*new(spec.Distribution),
//		*new(string),
//
//		// Slices of simple types.
//		*new([]int),
//		*new([]float64),
//		*new([]string),
//
//		// Slices of slices of simple types.
//		*new([][]int),
//		*new([][]float64),
//	}
//
//	// maxExamples represents the maximum number of inputsOutputs samples
//	// provided in a CLG profile. A CLG profile may contain only one sample in
//	// case the CLG interface is very strict. Nevertheless there might be CLGs
//	// that accept a variadic amount of input parameters or return a variadic
//	// amount of output results. The number of possible inputsOutputs samples can
//	// be infinite in theory. Thus we cap the amount of inputsOutputs samples by
//	// maxSamples.
//	maxSamples = 10
//
//	// numArgs is an ordered list of numbers used to find out how many input
//	// arguments a CLG expects. Usually CLGs do not expect more than 5 input
//	// arguments. For special cases we try to find out how many they expect
//	// beyond 5 arguments. Here we assume that a CLG might expect 10 or even 50
//	// arguments. In case a CLG expects 50 or more arguments, we assume it
//	// expects infinite arguments.
//	numArgs = []int{0, 1, 2, 3, 4, 5, 10, 20, 30, 40, 50}
//)
