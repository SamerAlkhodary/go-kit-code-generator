package main

import (
	"os"

	"github.com/samkhud/go-kit-code-generator/gen"
	"github.com/samkhud/go-kit-code-generator/parser"
)

func main() {
	args := os.Args[1:]

	parser := parser.Parser{}

	gen := gen.CreateGenerator(args[0], args[1], parser)
	gen.Generate()

}
