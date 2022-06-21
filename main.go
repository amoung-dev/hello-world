package main

import (
	"fmt"

	"github.com/amoung-dev/hello-world/hello"
	"github.com/amoung-dev/hello-world/world"
)

func main() {
	fmt.Println(hello.Greet())
	fmt.Println(world.Greet())
}
