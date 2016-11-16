package spec

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// PermutationService creates permutations of arbitrary lists as configured.
type PermutationService interface {
	Boot()
	Metadata() map[string]string
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
	PermuteBy(list objectspec.PermutationList, delta int) error
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
