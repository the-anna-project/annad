package shape

// stringShape represents a string shape permutation. A string can be permuted
// as follows.
//
//     This is how the initial string shape permutation looks like.
//
//         ""
//
//     This is how the second string shape permutation looks like.
//
//         "a"
//
//     This is how the third string shape permutation looks like.
//
//         "abc"
//
//     This is how the Nth string shape permutation looks like.
//
//         "hdebbbf"
//
type stringShape struct {
	// Index represents the index being incremented to identify the string shape
	// permutation.
	Index int

	// Indizes represents the Index's translation. Each rank of the Index is
	// represented separately within Indizes.
	//
	//    345 translates to []int{3, 4, 5}
	//
	Indizes []int

	// MaxLength represents the maximum length Member is allowed to grow.
	MaxLength int

	// Member represents the string shape being permuted. Initially this is the
	// zero value of string: "".
	Member string

	// Pieces represents the Member's translation. Each rang of the Member is
	// represented separately within Pieces.
	Pieces []string

	// Type represents the type of the string shape permutation: string.
	Type string

	// Values represents the values being used to permute Member.
	Values []string
}

func (ss *stringShape) GetIndex() int             {}
func (ss *stringShape) GetIndizes() []int         {}
func (ss *stringShape) GetMaxLength() int         {}
func (ss *stringShape) GetMember() interface{}    {}
func (ss *stringShape) GetPieces() []interface{}  {}
func (ss *stringShape) GetType() string           {}
func (ss *stringShape) GetValues() []interface{}  {}
func (ss *stringShape) PermuteBy(delta int) error {}
