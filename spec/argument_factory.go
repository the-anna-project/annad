package spec

// ArgumentFactory creates permutations of arguments as configured. It
// implements PermutationFactory. Anyway the following notes should be
// considered.
//
// PermuteBy permutes the given argument list using delta. Permutation is
// done in two dimensions.
//
// The first dimension of permutation applies to the whole argument list.
// Members of that list are permuted once. Then the second dimension of
// permutation is applied. After the second dimension was applied, the first
// one is applied and iteration starts from the beginning again.
//
// The second dimension of permutation applies to a single member of the
// argument list. Its members are permuted once. Then the next member's
// members are permuted, until all members are permuted once. Then the first
// dimension of permutation is applied and the iteration starts from the
// beginning again.
type ArgumentFactory interface {
	PermutationFactory
}
