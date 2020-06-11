package main

import (
	"go-kit-code-generator/gen"
	"go-kit-code-generator/parser"
	"os"
)

func main() {
	args := os.Args[1:]

	parser := parser.Parser{}

	gen := gen.CreateGenerator(args[0], args[1], parser)
	gen.Generate()

}
