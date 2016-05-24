package spec

// PermutationList is supposed to be permuted by a permutation factory.
type PermutationList interface {
	// GetIndizes returns the list's current indizes.
	GetIndizes() []int

	// GetMaxGrowth returns the list's configured growth limit. The growth limit
	// is used to prevent infinite permutations. E.g. MaxGrowth set to 4 will not
	// permute up to a list of 5 members.
	GetMaxGrowth() int

	// GetMembers returns the list's permuted members.
	GetMembers() []interface{}

	// GetValues returns the list's current values.
	GetValues() []interface{}

	// SetIndizes sets the given indizes of the current permutation list.
	SetIndizes(indizes []int)

	// SetMembers sets the given members of the current permutation list.
	SetMembers(members []interface{})
}
