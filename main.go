package main

import (
	"fmt"
	"io"
	"os"

	"github.com/amoung-dev/hello-world/hello"
	"github.com/amoung-dev/hello-world/world"
)

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
}

func main() {
	displayGreetings(os.Stdout)
}
