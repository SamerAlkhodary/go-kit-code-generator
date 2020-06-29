package gen

import (
	"fmt"
	"log"
	"os"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

type repositroyGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createRepositoryGenerator(s model.Service, outputFile string) fileGenerator {
	return &repositroyGenerator{
		outputFile: outputFile,
		s:          s,
	}
}

func (rg *repositroyGenerator) run(outputPath string) {
	rg.generateCode()
	rg.generateFile(outputPath)
}
func (rg *repositroyGenerator) generateCode() {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s", rg.s.GetServiceName())
	fmt.Fprintf(&code, "\nimport(\n%q\n%q\n%q\n)", "context", "github.com/go-kit/kit/log", "database/sql")
	fmt.Fprintf(&code, " \ntype Repository interface{\n")
	for _, endpoint := range rg.s.Endpoints {
		fmt.Fprintf(&code, "\n%s(", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s %s,", rg.s.GetVariableName(arg, true), rg.s.GetType(arg))

		}
		fmt.Fprintf(&code, "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", rg.s.GetType(out))

		}
		fmt.Fprintf(&code, "error)")

	}
	fmt.Fprintf(&code, "\n}\n")
	fmt.Fprintf(&code, "type repository struct{\n db *sql.DB\nlogger log.Logger\n}\n")
	fmt.Fprintf(&code, "\nfunc MakeNewRepository(db *sql.DB, logger log.Logger)Repository{\n ")
	fmt.Fprintf(&code, "\nreturn &repository{\ndb:db,\nlogger:log.With(logger,%q,%q ),\n}\n", "repository", "sql")
	fmt.Fprintf(&code, "\n}\n")
	for _, endpoint := range rg.s.Endpoints {
		fmt.Fprintf(&code, "\nfunc(repo *repository)%s(", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s %s,", rg.s.GetVariableName(arg, true), rg.s.GetType(arg))
		}
		fmt.Fprintf(&code, "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", rg.s.GetType(out))
		}
		fmt.Fprintf(&code, "error){\n}\n")

	}

	rg.code = code.String()
}
func (rg *repositroyGenerator) generateFile(outputPath string) {
	if !rg.s.Repository.Value {
		log.Println(("no repository required"))
		return
	}

	var path string

	path = fmt.Sprintf("%s/%s.go", outputPath, rg.outputFile)

	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(rg.code)
	defer file.Close()

}
