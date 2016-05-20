// Package permutation provides a simple permutation factory implementation in
// which the order of the members of a permutation list is permuted. Advantages
// of the permutation factory is memory effiency and reproducability. It is
// memory efficient because all possible combinations are not stored in memory,
// but created on demand. Depending on the provided delta this can be quiet
// fast in case it is not too big. The factory is reproducible because of the
// index used to represent a permutation.
//
//     This is how the initial factory permutation looks like.
//
//         []interface{"a"}
//
//     This is how the second factory permutation looks like.
//
//         []interface{"a", 7}
//
//     This is how the third factory permutation looks like.
//
//         []interface{"a", 7, []float64{2.88}}
//
//     This is how the Nth factory permutation looks like.
//
//         []interface{7, []float64{2.88}, "a"}
//
package permutation
