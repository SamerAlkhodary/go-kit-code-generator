package gen

import (
	"fmt"
	"go-kit-code-generator/model"
	"go-kit-code-generator/parser"
	"log"
	"os"
	"sync"
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
	var wg sync.WaitGroup
	log.Println("Generating :", gen.inputPath)
	service := gen.parser.Parse(gen.inputPath)
	service.Apply()
	err := service.CheckForError()
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	endCode := endpointsGenerator(*service)
	serviceCode := serviceGenerator(*service)
	transportCode := transportGenerator(*service)
	encodersCode := encoderDecoderGenerator(*service)
	modelCode := modelGenerator(*service)
	serverCode := mainCodeGenerator(*service)
	repoCode := repositroyGenerator(*service)
	wg.Add(7)
	go genCode(service, gen, "endpoints", endCode, &wg)
	go genCode(service, gen, "transport", transportCode, &wg)
	go genCode(service, gen, "encoders", encodersCode, &wg)
	go genCode(service, gen, "model", modelCode, &wg)
	go genCode(service, gen, "server", serverCode, &wg)
	go genCode(service, gen, "service", serviceCode, &wg)
	go genCode(service, gen, "repository", repoCode, &wg)
	wg.Wait()

	fmt.Println("Generating acomplished")
}
func genCode(s *model.Service, gen *generator, name string, code string, wg *sync.WaitGroup) {
	defer wg.Done()
	if name == "repository" {
		fmt.Println(s.Repository)
		if !s.Repository {
			return
		}
	}

	path := fmt.Sprintf("%s/%s.go", gen.outputPath, name)
	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(code)
	defer file.Close()

}
