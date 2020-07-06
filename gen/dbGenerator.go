package gen

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/samkhud/go-kit-code-generator/model"
)

type dbGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createDbGenerator(s model.Service, outputFile string) fileGenerator {
	return &dbGenerator{
		outputFile: outputFile,
		s:          s,
	}
}
func (dg *dbGenerator) GetFileName() string {
	return dg.outputFile
}

func (dg *dbGenerator) run(outputPath string) {
	dg.generateCode()
	dg.generateFile(outputPath)
}

type Pair struct {
	modelName string
	primeKey  model.Attribute
}

func (dg *dbGenerator) generateCode() {
	var arrs []model.Attribute

	var code strings.Builder
	code.Grow(1000)
	dependencies := make(map[string][]Pair)
	for _, m := range dg.s.Models {
		fmt.Fprintf(&code, "\nCREATE TABLE %ss\n(\n", m.GetName(true))

		for _, attr := range m.GetModelAttributes() {
			if dg.s.IsArray(attr.DataType) {
				arrs = append(arrs, attr)
			}
			if !dg.s.IsAddedType(attr.DataType) && !dg.s.IsArray(attr.DataType) {
				fmt.Fprintf(&code, "%s %s ,\n", attr.GetName(true), attr.GetDBType())
			} else {

				p := Pair{
					modelName: m.GetName(true),
					primeKey:  findPrimaryKey(m.GetModelAttributes()),
				}

				dependencies[attr.GetName(true)] = append(dependencies[attr.GetName(true)], p)
				log.Println(dependencies)

			}
		}
		deps := dependencies[m.GetName(true)]
		for _, forgen := range deps {
			fmt.Fprintf(&code, "%s%s %s  NOT NULL  ,\n", forgen.modelName, forgen.primeKey.GetName(false), "--TODO write type")
		}

		fmt.Fprintf(&code, "timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP")
		fmt.Fprintf(&code, "\n);")
	}
	for _, arrElement := range arrs {
		fmt.Fprintf(&code, "%s", arrayHandler(arrElement, dependencies[arrElement.GetName(true)][0]))
	}
	dg.code = code.String()

}
func findPrimaryKey(attrs []model.Attribute) model.Attribute {
	for _, attr := range attrs {
		if attr.IsPrimaryKey() {
			return attr
		}

	}
	return model.Attribute{}

}
func (dg *dbGenerator) generateFile(outputPath string) {
	if dg.s.Repository.DB.Name != "mysql" && dg.s.Repository.DB.Name != "postgress" {
		return
	}
	var path string
	path = fmt.Sprintf("%s/%s.sql", outputPath, dg.outputFile)
	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(dg.code)
	defer file.Close()
}
func arrayHandler(attr model.Attribute, p Pair) string {
	var tmp strings.Builder
	fmt.Fprintf(&tmp, "\nCREATE TABLE %s\n(\n", attr.GetName(true))
	fmt.Fprintf(&tmp, "%s%s %s  NOT NULL  ,\n", p.modelName, p.primeKey.GetName(true), "--TODO write type")
	fmt.Fprintf(&tmp, "%s %s   NULL  ,\n", "element", attr.GetDBType())
	fmt.Fprintf(&tmp, "timestamp   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP")
	fmt.Fprintf(&tmp, "\n);")
	return tmp.String()

}
