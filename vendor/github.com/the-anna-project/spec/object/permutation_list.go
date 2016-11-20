package spec

// PermutationList is supposed to be permuted by a permutation service.
type PermutationList interface {
	// Indizes returns the list's current indizes.
	Indizes() []int
	// MaxGrowth returns the list's configured growth limit. The growth limit
	// is used to prevent infinite permutations. E.g. MaxGrowth set to 4 will not
	// permute up to a list of 5 raw values.
	MaxGrowth() int
	// PermutedValues returns the list's permuted values.
	PermutedValues() []interface{}
	// RawValues returns the list's unpermuted raw values.
	RawValues() []interface{}
	// SetIndizes sets the given indizes of the current permutation list.
	SetIndizes(indizes []int)
	SetMaxGrowth(maxGrowth int)
	SetRawValues(rawValues []interface{})
}
