// Package shape provides a special form of permutation factory in which not
// the order of members is permuted but the shape of the members itself. This
// can be used for e.g. CLG arguments creation. The following shows an example
// of a shape factory permutation.
//
//     This is how the initial shape factory permutation looks like.
//
//         []interface{"a", 0, []float64{0.0}}
//
//     This is how the second shape factory permutation looks like.
//
//         []interface{"a", 0, []float64{0.1}}
//
//     This is how the third shape factory permutation looks like.
//
//         []interface{"a", 0, []float64{0.2}}
//
//     This is how the Nth shape factory permutation looks like.
//
//         []interface{"abcd", 4, []float64{0.3}}
//
package shape

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

	// MaxGrowth represents the maximum each Member is allowed to grow.
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
