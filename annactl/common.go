package main

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
