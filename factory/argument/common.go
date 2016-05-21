package argument

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func shiftPointer(pointer, base int) (int, bool) {
	pointer++
	var reset bool

	if pointer >= base {
		pointer = 0
		reset = true
	}

	return pointer, reset
}
