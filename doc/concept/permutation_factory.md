# permutation factory
The permutation factory permutes the order of the members of an arbitrary list.
Advantages of the permutation factory is memory effiency and reproducability.
It is memory efficient because all possible combinations are not stored in
memory, but created on demand. Depending on the provided delta the creation can
be quiet fast in case the delta is not too big. The factory is reproducible
because of the index used to represent a permutation. So in case the given
delta is way too big one might want to provide the indizes directly. Then the
permutation is pretty fast because it is basically about looking up some map
entries.

### example

Imagine the following example.

```
[]interface{"a", 7, []float64{2.88}}
```

This is how the initial factory permutation looks like. In fact, there is no
permutation.

```
[]interface{}
```

This is how the first factory permutation looks like.

```
[]interface{"a"}
```

This is how the second factory permutation looks like.

```
[]interface{7}
```

This is how the third factory permutation looks like.

```
[]interface{[]float64{2.88}}
```

This is how the Nth factory permutation looks like.

```
[]interface{[]float64{2.88}, "a"}
```

### reference

See the following code documentation for more information.

- https://godoc.org/github.com/xh3b4sd/anna/factory/permutation
- https://godoc.org/github.com/xh3b4sd/anna/spec#PermutationFactory
