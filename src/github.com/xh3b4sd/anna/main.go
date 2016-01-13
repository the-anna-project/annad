package main

import (
	"fmt"

	"github.com/xh3b4sd/anna/server"
)

func main() {
	fmt.Printf("%#v\n", "Hello, I am Anna.")

	fmt.Printf("%#v\n", "starting server")
	server.Listen()
}
