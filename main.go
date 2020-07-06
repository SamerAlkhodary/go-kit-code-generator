package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/samkhud/go-kit-code-generator/gen"
	"github.com/samkhud/go-kit-code-generator/parser"
)

func handleError(msg string) {
	log.Fatalf("Rutime error:%s\n%s", msg, "Run --help flag to get some help")
}

func setup(exec string) string {
	var code strings.Builder
	fmt.Fprintf(&code, "%s is a tool to generate go code.", exec)
	fmt.Fprintf(&code, "\nUsage:")
	fmt.Fprintf(&code, "\n\t%s <command> [arguments]", exec)
	fmt.Fprintf(&code, "\nThe commands are:")
	fmt.Fprintf(&code, "\n\tgt\tgenerates an empty .yaml template file")
	fmt.Fprintf(&code, "\n\tgs\tgenerates a service based on a .yaml file")
	s := fmt.Sprintf("%s help <command>", exec)
	fmt.Fprintf(&code, "\nUse %q for more information about a command.", s)
	return code.String()

}
func runCLI() {
	parser := parser.Parser{}

	args := os.Args[0:]
	code := setup(os.Args[0])
	if len(args) < 2 {
		fmt.Println(code)
		return

	}
	gen := gen.CreateGenerator(parser)
	switch strings.ToLower(args[1]) {
	case "gs":
		if len(args) < 3 {
			handleError(" gs: missing path to .yaml file")
		}
		if len(args) < 4 {
			handleError(" gs: missing output destination")
		}

		gen.GenerateService(args[2], args[3])
		break
	case "gt":
		if len(args) < 3 {
			handleError(" gt: missing output path")

		}
		if len(args) < 4 {
			handleError(" gt: missing yaml file name")

		}
		gen.GenerateTemplate(args[2], args[3])
		break
	case "help":
		if len(args) < 3 {
			fmt.Println(code)
			return

		}
		if len(args) == 3 {
			switch args[2] {
			case "gs":
				fmt.Printf("usage: %s gs [.yaml path] [output path] \n", args[0])

				break
			case "gt":
				fmt.Printf("usage: %s gt [output path] [yaml file name]\n", args[0])
				break
			default:
				fmt.Println("unrecognised command")
				return

			}

		}

		break
	default:
		handleError(" Unrecognised action flag ")
		break

	}
}
func main() {
	runCLI()

}