package spec

// PermutationList is supposed to be permuted by a permutation service.
type PermutationList interface {
	// GetIndizes returns the list's current indizes.
	GetIndizes() []int

	// GetMaxGrowth returns the list's configured growth limit. The growth limit
	// is used to prevent infinite permutations. E.g. MaxGrowth set to 4 will not
	// permute up to a list of 5 raw values.
	GetMaxGrowth() int

	// GetPermutedValues returns the list's permuted values.
	GetPermutedValues() []interface{}

	// GetRawValues returns the list's unpermuted raw values.
	GetRawValues() []interface{}

	// SetIndizes sets the given indizes of the current permutation list.
	SetIndizes(indizes []int)
}
