package gen

import (
	"fmt"
	"time"

	"github.com/samkhud/go-kit-code-generator/model"

	"log"
	"os"
	"sync"

	"github.com/samkhud/go-kit-code-generator/parser"
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

	wg.Add(8)
	start := time.Now()
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		endCode := endpointsGenerator(*service)
		genCode(service, gen, "endpoints", endCode)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		transportCode := transportGenerator(*service)
		genCode(service, gen, "transport", transportCode)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		encoderCode := encoderDecoderGenerator(*service)
		genCode(service, gen, "encoders", encoderCode)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		modelCode := modelGenerator(*service)
		genCode(service, gen, "model", modelCode)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		serverCode := mainCodeGenerator(*service)
		genCode(service, gen, "server", serverCode)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		serviceCode := serviceGenerator(*service)
		genCode(service, gen, "service", serviceCode)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		repoCode := repositroyGenerator(*service)
		genCode(service, gen, "repository", repoCode)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		cacheCode := cacheGenerator(*service)
		genCode(service, gen, "cache", cacheCode)

	}(&wg)

	wg.Wait()
	fmt.Println(time.Since(start))

	fmt.Println("Generating acomplished")
}
func genCode(s *model.Service, gen *generator, name string, code string) {
	if name == "repository" {
		if !s.Repository.Value {
			return
		}
	}
	if name == "cache" {
		if s.RedisCache.GetHost() == "" {
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
