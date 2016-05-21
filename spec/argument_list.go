package spec

// ArgumentList represents a list of permutation lists.
type ArgumentList interface {
	// CreateArguments converts the list's members to their underlying concrete
	// type. The order of the argument list returned reflects the order of the
	// lists permuted members. There might be an error returned in case an
	// unsupported concrete type was detected.
	CreateArguments() ([]interface{}, error)

	// GetPointer returns the list's pointer. It describes the member to permute
	// next.
	GetPointer() int

	// SetPointer sets the given pointer to the current argument list.
	SetPointer(pointer int)

	PermutationList
}
