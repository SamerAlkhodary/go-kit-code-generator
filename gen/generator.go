package gen

import (
	"fmt"
	"go-kit-code-generator/parser"
	"log"
	"os"
)

type generator struct {
	inputPath  string
	outputPath string
	parser     parser.Parser
}

func CreateGenerator(inputPath string, outputPath string, parser parser.Parser) *generator {
	return &generator{
		inputPath:  inputPath,
		outputPath: outputPath,
		parser:     parser,
	}
}
func (gen *generator) GetInputPath() string {
	return gen.inputPath
}

func (gen *generator) GetOutputPath() string {
	return gen.outputPath
}
func (gen *generator) GetParser() *parser.Parser {
	return &gen.parser
}
func (gen *generator) Generate() {
	log.Println("Generating :", gen.inputPath)
	service := gen.parser.Parse(gen.inputPath)
	endCode := endpointsGenerator(*service)
	serviceCode := serviceGenerator(*service)
	transportCode := transportGenerator(*service)
	encodersCode := encoderDecoderGenerator(*service)
	mainCode := mainCodeGenerator(*service)
	endpath := fmt.Sprintf("%s/endpoints.go", gen.outputPath)
	servicepath := fmt.Sprintf("%s/service.go", gen.outputPath)
	transportPath := fmt.Sprintf("%s/transport.go", gen.outputPath)
	encodersPath := fmt.Sprintf("%s/encoders.go", gen.outputPath)
	mainPath := fmt.Sprintf("%s/main.go", gen.outputPath)
	mainFile, err := os.Create(mainPath)
	encodersFile, err := os.Create(encodersPath)
	endFile, err := os.Create(endpath)
	serviceFile, err := os.Create(servicepath)
	transportFile, err := os.Create(transportPath)

	if err != nil {
		log.Printf("error while creating file:%v", err)
	}
	defer mainFile.Close()
	defer transportFile.Close()
	defer serviceFile.Close()
	defer encodersFile.Close()
	defer endFile.Close()
	transportFile.WriteString(transportCode)
	endFile.WriteString(endCode)
	serviceFile.WriteString(serviceCode)
	encodersFile.WriteString(encodersCode)
	mainFile.WriteString(mainCode)
}
