package spec

// PermutationFactory creates permutations of arbitrary lists as configured and
// implemented.
type PermutationFactory interface {
	// CreateIndexWithIndizes takes a list of indizes and translates it to an
	// index representation. That is a string.
	//
	//    []int{3, 4, 5} translates to  345
	//
	CreateIndexWithIndizes(indizes []int) string

	// CreateMembersWithIndizes takes a list of indizes and translates it to a
	// permuted list based on the configured values.
	CreateMembersWithIndizes(indizes []int) []interface{}

	// GetIndex returns the factory's current index.
	GetIndex() string

	// GetIndizes returns the factory's current indizes.
	GetIndizes() []int

	// MaxGrowth returns the factory's configured growth limit. The growth limit
	// is used to prevent infinite permutations. E.g. MaxGrowth set to 4 will not
	// permute up to a list of 5 members.
	GetMaxGrowth() int

	// GetMembers returns the factory's current members.
	GetMembers() []interface{}

	// GetValues returns the factory's current values.
	GetValues() []interface{}

	// PermuteBy permutes the configured values by applying the given delta to
	// the currently configured indizes. Error might indicate that the configured
	// max growth is reached.
	PermuteBy(delta int) error

	// UpdateIndizesWithDelta processes the essentiaal operation of permuting the
	// factory's indizes using the given delta. Indizes are capped by the amount
	// of configured values. E.g. having a list of two values configured will
	// increment the indizes in a base 2 number system.
	//
	//     Imagine the following values being configured.
	//
	//         []interface{"a", "b"}
	//
	//     Here the index 0 translates to indizes []int{0} and causes the
	//     following permutation.
	//
	//         []interface{"a"}
	//
	//     Here the index 1 translates to indizes []int{1} and causes the
	//     following permutation.
	//
	//         []interface{"b"}
	//
	//     Here the index 00 translates to indizes []int{0, 0} and causes the
	//     following permutation.
	//
	//         []interface{"a", "a"}
	//
	UpdateIndizesWithDelta(delta int) error
}
