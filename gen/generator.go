package gen

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

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
	var generators []fileGenerator
	generators = append(generators, createCacheGenerator(*service, "cache"))
	generators = append(generators, createDbGenerator(*service, "db"))
	generators = append(generators, createEncodersGenerator(*service, "encoders"))
	generators = append(generators, createEndpointGenerator(*service, "endpoints"))
	generators = append(generators, createServiceGenerator(*service, "service"))
	generators = append(generators, createTransportGenerator(*service, "transport"))
	generators = append(generators, createRepositoryGenerator(*service, "repository"))
	generators = append(generators, createMainGenerator(*service, "facade"))
	generators = append(generators, createModelGenerator(*service, "model"))

	createPath(outputPath)
	start := time.Now()
	for _, generator := range generators {
		wg.Add(1)
		func() {

			defer wg.Done()
			generator.run(gen.outputPath)

		}()

	}

	wg.Wait()
	fmt.Println(time.Since(start))

	fmt.Println("Generating acomplished")
}
