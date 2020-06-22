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

func CreateGenerator(parser parser.Parser) *generator {
	return &generator{

		parser: parser,
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
func (gen *generator) GenerateTemplate(path string, title string) {
	code := templateGenerator()
	createPath(path)
	p := fmt.Sprintf("%s/%s", path, title)

	file, err := os.Create(p)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(code)
	defer file.Close()

}
func (gen *generator) GenerateDockerImage() {
	log.Println("Generating :", gen.inputPath)
	service, er := gen.parser.Parse(gen.inputPath)
	if er != nil {
		log.Fatal(er)
	}
	service.Apply()
	err := service.CheckForError()
	if err != nil {
		log.Fatal(err)
	}

}
func createPath(p string) {
	e := os.Mkdir(p, 0700)
	if e != nil {
		log.Printf("error:%v", e)
	}

}
func (gen *generator) GenerateService(inputPath string, outputPath string) {
	var wg sync.WaitGroup
	gen.inputPath = inputPath
	gen.outputPath = outputPath

	log.Println("Generating :", gen.inputPath)
	service, er := gen.parser.Parse(gen.inputPath)
	if er != nil {
		log.Fatal(er)
	}
	service.Apply()
	err := service.CheckForError()
	if err != nil {
		log.Fatal(err)
	}
	createPath(outputPath)
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
		log.Println(s.Repository)
		if !s.Repository.Value {
			log.Println(("no repository required"))
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
