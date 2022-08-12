package main

import (
	"github.com/praveensanap/glox-interpreter/compiler"
	"os"
)

func main() {
	args := os.Args[1:]
	file := ""
	if len(args) > 0 {
		file = args[0]
	}
	compiler.Compile(file)
}
