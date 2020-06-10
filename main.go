package main

import (
	"os"

	"services/generator/gen"
	"services/generator/parser"
)

func main() {
	args := os.Args[1:]

	parser := parser.Parser{}

	gen := gen.CreateGenerator(args[0], args[1], parser)
	gen.Generate()

}
