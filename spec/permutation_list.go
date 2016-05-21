package spec

// PermutationType describes the type that all joined members represent. A
// permutation might be done to permute a concrete type. This concrete type is
// then used to create a PermutationList. Here the concrete type represents the
// PermutationType. Once permuted the member's of the list can be converted to
// the concrete type.
//
// Note that the permutation type might be empty in case it is irrelevant for
// the implementation.
//
type PermutationType string

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

	// GetType returns the list's permutation type.
	GetType() PermutationType

	// GetValues returns the list's current values.
	GetValues() []interface{}

	// SetIndizes sets the given indizes of the current permutation list.
	SetIndizes(indizes []int)

	// SetMembers sets the given members of the current permutation list.
	SetMembers(members []interface{})
}
