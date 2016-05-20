package spec

// PermutationList is supposed to be permuted by a permutation factory.
type PermutationList interface {
	// GetIndizes returns the factory's current indizes.
	GetIndizes() []int

	// GetMaxGrowth returns the factory's configured growth limit. The growth limit
	// is used to prevent infinite permutations. E.g. MaxGrowth set to 4 will not
	// permute up to a list of 5 members.
	GetMaxGrowth() int

	// GetMembers returns the factory's permuted members.
	GetMembers() []interface{}

	// GetValues returns the factory's current values.
	GetValues() []interface{}

	// SetIndizes sets the given indizes of the current permutation list.
	SetIndizes(indizes []int)

	// SetMembers sets the given members of the current permutation list.
	SetMembers(members []interface{})
}
