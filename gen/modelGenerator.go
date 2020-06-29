package gen

import (
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/samkhud/go-kit-code-generator/model"
)

type modelGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createModelGenerator(s model.Service, outputFile string) fileGenerator {
	return &modelGenerator{
		outputFile: outputFile,
		s:          s,
	}
}
func (mg *modelGenerator) run(outputPath string) {
	mg.generateCode()
	mg.generateFile(outputPath)
}

func (mg *modelGenerator) generateCode() {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s", mg.s.GetServiceName())
	for _, model := range mg.s.Models {
		fmt.Fprintf(&code, "\ntype %s struct{\n", model.GetName(false))

		for _, attr := range model.GetModelAttributes() {

			fmt.Fprintf(&code, "\n%s %s `json:%q`\n", attr.GetName(false), attr.GetType(), attr.GetName(true))
		}
		fmt.Fprintf(&code, "}")
		fmt.Fprintf(&code, "\nfunc MakeNew%s(", model.GetName(false))
		for i, attr := range model.GetModelAttributes() {

			fmt.Fprintf(&code, "%s %s ", attr.GetName(true), attr.GetType())
			if i < len(model.GetModelAttributes())-1 {
				fmt.Fprintf(&code, ",")
			}
		}
		fmt.Fprintf(&code, ")*%s{\n return &%s{\n", model.GetName(false), model.GetName(false))
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\n%s : %s, \n", attr.GetName(false), attr.GetName(true))
		}
		fmt.Fprintf(&code, "}\n}")
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\nfunc(%s *%s)Get%s()%s{\nreturn %s.%s\n} \n",
				model.GetName(true),
				model.GetName(false),
				attr.GetName(false),
				attr.GetType(),
				model.GetName(true),
				attr.GetName(false))
			fmt.Fprintf(&code, "\nfunc(%s *%s)Set%s(new%s %s){\n%s.%s=new%s\n} \n",
				model.GetName(true),
				model.GetName(false),
				attr.GetName(false),
				attr.GetName(false),
				attr.GetType(),
				model.GetName(true),
				attr.GetName(false),
				attr.GetName(false))

		}

	}
	mg.code = code.String()
}
func (mg *modelGenerator) generateFile(outputPath string) {
	var path string

	path = fmt.Sprintf("%s/%s.go", outputPath, mg.outputFile)

	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(mg.code)
	defer file.Close()

}
