package gen

import (
	"fmt"
	"log"
	"os"
	"strings"

	"services/generator/model"
	"services/generator/parser"
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
	endpath := fmt.Sprintf("%s/endpoints.go", gen.outputPath)
	servicepath := fmt.Sprintf("%s/service.go", gen.outputPath)
	transportPath := fmt.Sprintf("%s/transport.go", gen.outputPath)

	endFile, err := os.Create(endpath)
	serviceFile, err := os.Create(servicepath)
	transportFile, err := os.Create(transportPath)

	if err != nil {
		log.Printf("error while creating file:%v", err)
	}
	defer transportFile.Close()
	defer serviceFile.Close()
	defer endFile.Close()
	transportFile.WriteString(transportCode)
	endFile.WriteString(endCode)
	serviceFile.WriteString(serviceCode)

	//TODO: generates code and writes to file

}
func transportGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", serviceName(s.Name))
	fmt.Fprintf(&code, "import(httptransport %q\n%q\n%q)\n", "github.com/go-kit/kit/transport/http", "context", "github.com/gorilla/mux")
	fmt.Fprintf(&code, "func NewHTTPServer(ctx context.Context,endPoints Endpoints)http.Handler{\n")
	fmt.Fprintf(&code, "r:=mux.NewRouter()\nr.Use(commonMiddleware)\n")
	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "r.Method(/*TODO:choose request*/).Path(/*TODO: choose path*/).Handler(httptransport.NewServer(\nendpoints.%s,\ndecode%sRequest,\nencode%sResponse,\n))\n", endpoint.Name, endpoint.Name, endpoint.Name)

	}
	fmt.Fprintf(&code, "\nreturn r")

	fmt.Fprintf(&code, "\n}\n")
	return code.String()
}
func endpointsGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", serviceName(s.Name))
	fmt.Fprintf(&code, "import(%q)\n", "github.com/go-kit/kit/endpoint")

	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "type %sRequest struct{\n", endpoint.Name)
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s\n", arg)

		}
		fmt.Fprintf(&code, "%s\n}", "")
		fmt.Fprintf(&code, "\ntype  %sResponse struct{\n", endpoint.Name)
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s\n", out)

		}
		fmt.Fprintf(&code, "\n}\n func make%sEndpoint(s Service)endpoint.Endpoint{\nreturn func(ctx context.Context, request interface{}) (interface{}, error) {\nreq := request.(%sRequest)\n", endpoint.Name, endpoint.Name)
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", strings.Split(out, " ")[0])
		}
		fmt.Fprintf(&code, "error:=s.%s(", endpoint.Name)

		for i, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "req.%s", strings.Split(strings.TrimSpace(arg), " ")[0])
			if i < len(endpoint.GetArgs())-1 {
				fmt.Fprintf(&code, ",")

			}
		}
		fmt.Fprintf(&code, ")\nreturn %sResponse{", endpoint.Name)
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s: %s,", strings.Split(out, " ")[0], strings.Split(out, " ")[0])
		}
		fmt.Fprintf(&code, "}, error\n}\n}\n")
	}
	fmt.Fprintf(&code, "%s", "type Endpoints struct{\n ")
	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "%s endpoint.Endpoint\n", endpoint.Name)

	}
	fmt.Fprintf(&code, "%s\n", "}")

	fmt.Fprintf(&code, "func MakeEndpoints(s %s)Endpoints{\n return Endpoints{\n", s.Name)

	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "%s:make%s(s),\n", endpoint.Name, endpoint.Name)

	}

	fmt.Fprintf(&code, "%s\n}", "}")

	return code.String()
}
func serviceGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", serviceName(s.Name))
	fmt.Fprintf(&code, "import(%q\n %q)\n", "context", "github.com/go-kit/kit/log")

	fmt.Fprintf(&code, "type %s interface{\n", serviceName(s.Name))
	for _, endpoint := range s.Endpoints {
		fmt.Fprintf(&code, "%s(", endpoint.Name)
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s,", arg)
		}
		fmt.Fprintf(&code, "%s", "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Println(strings.Split(out, " ")[1])
			fmt.Fprintf(&code, "%s,", strings.Split(out, " ")[1])
		}
		fmt.Fprintf(&code, "error)\n")

	}
	fmt.Fprintf(&code, "%s\n", "}")

	fmt.Fprintf(&code, "type %s struct{\n", serviceName(s.Name))
	fmt.Fprintf(&code, "%s\n", "repository Repository")
	fmt.Fprintf(&code, "%s\n", "logger log.Logger")

	fmt.Fprintf(&code, "%s\n", "}")
	fmt.Fprintf(&code, "func NewService(rep Repository,logger log.Logger)%s{\n return &%s{\n repository: rep,\n logger:logger,\n}}\n", serviceName(s.Name), serviceName(s.Name))
	for _, endpoint := range s.Endpoints {
		fmt.Fprintf(&code, "func(s *%s)%s(", s.Name, endpoint.Name)
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s,", arg)
		}
		fmt.Fprintf(&code, "%s", "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Println(strings.Split(out, " ")[1])
			fmt.Fprintf(&code, "%s,", strings.Split(out, " ")[1])
		}
		fmt.Fprintf(&code, "error){\n")
		fmt.Fprintf(&code, "Logger:= log.With(s.logger,%q,%s)\n//TODO: implement\n", "method", endpoint.Name)

		fmt.Fprintf(&code, "}\n")

	}
	return code.String()
}
func serviceName(name string) string {
	words := strings.Split(strings.ToLower(name), "")
	res := ""
	for _, word := range words {
		res = fmt.Sprintf("%s%s", res, strings.TrimSpace(word))
	}
	return res

}
