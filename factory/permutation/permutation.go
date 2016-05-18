// Package permutatation provides a simple permutation factory in which the
// order of the members of a list is permuted.
//
//     This is how the initial factory permutation looks like.
//
//         []interface{"a"}
//
//     This is how the second factory permutation looks like.
//
//         []interface{"a", 7}
//
//     This is how the third factory permutation looks like.
//
//         []interface{"a", 7, []float64{2.88}}
//
//     This is how the Nth factory permutation looks like.
//
//         []interface{7, []float64{2.88}, "a"}
//
package permutatation

// TODO
var (
	// argTypes represents a list of well known types used to identify CLG input-
	// and output types. Here we want to have a list of types only.
	//
	// TODO identify types by strings
	argTypes = []interface{}{
		// Simple types.
		*new(bool),
		*new(int),
		*new(float64),
		*new(string),
		*new(spec.Distribution),
		*new(string),

		// Slices of simple types.
		*new([]int),
		*new([]float64),
		*new([]string),

		// Slices of slices of simple types.
		*new([][]int),
		*new([][]float64),
	}

	// maxExamples represents the maximum number of inputsOutputs samples
	// provided in a CLG profile. A CLG profile may contain only one sample in
	// case the CLG interface is very strict. Nevertheless there might be CLGs
	// that accept a variadic amount of input parameters or return a variadic
	// amount of output results. The number of possible inputsOutputs samples can
	// be infinite in theory. Thus we cap the amount of inputsOutputs samples by
	// maxSamples.
	maxSamples = 10

	// numArgs is an ordered list of numbers used to find out how many input
	// arguments a CLG expects. Usually CLGs do not expect more than 5 input
	// arguments. For special cases we try to find out how many they expect
	// beyond 5 arguments. Here we assume that a CLG might expect 10 or even 50
	// arguments. In case a CLG expects 50 or more arguments, we assume it
	// expects infinite arguments.
	numArgs = []int{0, 1, 2, 3, 4, 5, 10, 20, 30, 40, 50}
)

type factory struct {
	// Index represents the index being incremented to identify the string shape
	// permutation.
	Index int

	// Indizes represents the Index's translation. Each rank of the Index is
	// represented separately within Indizes.
	//
	//    345 translates to []int{3, 4, 5}
	//
	Indizes []int

	// MaxGrowth represents the maximum length Members is allowed to grow.
	MaxGrowth int

	// Members represents the list being permuted. Initially this is the zero
	// value of []interface{}: []interface{}{}.
	Members []interface{}

	// Values represents the values being used to permute Members.
	Values []interface{}
}

func (f *factory) GetIndex() int             {}
func (f *factory) GetIndizes() []int         {}
func (f *factory) GetMaxGrowth() int         {}
func (f *factory) GetMembers() []interface{} {}
func (f *factory) GetValues() []interface{}  {}
func (f *factory) PermuteBy(delta int) error {}
