package spec

// FactoryCollection represents a collection of factories. This scopes different
// factory implementations in a simple container, which can easily be passed
// around.
type FactoryCollection interface {
	// ID represents an ID factory. It is used to create IDs of a certain type.
	ID() IDFactory

	// Permutation represents a permutation factory. It is used to permute
	// instances of type PermutationList.
	Permutation() PermutationFactory

	// Random represents a random factory. It is used to create random numbers.
	Random() RandomFactory
}

// FactoryProvider should be implemented by every object which wants to use
// factories. This then creates an API between factory implementations and
// factory users.
type FactoryProvider interface {
	Factory() FactoryCollection
}

// IDType represents some kind of configuration for ID creation.
type IDType int

// IDFactory creates pseudo random hash generation used for ID assignment.
type IDFactory interface {
	// New tries to create a new object ID using the configured ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	New() (ObjectID, error)

	// WithType tries to create a new object ID using the given ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	WithType(idType IDType) (ObjectID, error)
}

// PermutationFactory creates permutations of arbitrary lists as configured.
type PermutationFactory interface {
	// PermuteBy permutes the configured values by applying the given delta to
	// the currently configured indizes. Error might indicate that the configured
	// max growth is reached. It processes the essentiaal operation of permuting
	// the list's indizes using the given delta. Indizes are capped by the
	// amount of configured values. E.g. having a list of two values configured
	// will increment the indizes in a base 2 number system.
	//
	//     Imagine the following values being configured.
	//
	//         []interface{"a", "b"}
	//
	//     Here the index 0 translates to indizes []int{0} and causes the
	//     following permutation on the raw values.
	//
	//         []interface{"a"}
	//
	//     Here the index 1 translates to indizes []int{1} and causes the
	//     following permutation on the raw values.
	//
	//         []interface{"b"}
	//
	//     Here the index 00 translates to indizes []int{0, 0} and causes the
	//     following permutation on the raw values.
	//
	//         []interface{"a", "a"}
	//
	PermuteBy(list PermutationList, delta int) error
}

// PermutationList is supposed to be permuted by a permutation factory.
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

// RandomFactory creates pseudo random numbers. The factory might implement
// retries using backoff strategies and timeouts.
// TODO provide CreateMax to only return one random number
type RandomFactory interface {
	// CreateNMax tries to create a list of new pseudo random numbers. n
	// represents the number of pseudo random numbers in the returned list. The
	// generated numbers are within the range [0 max).
	CreateNMax(n, max int) ([]int, error)
}
