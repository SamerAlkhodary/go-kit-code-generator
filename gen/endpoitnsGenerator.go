package gen

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/samkhud/go-kit-code-generator/model"
)

type endpointsGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func (eg *endpointsGenerator) run(outputPath string) {
	eg.generateCode()
	eg.generateFile(outputPath)
}
func createEndpointGenerator(s model.Service, outputFile string) fileGenerator {
	return &endpointsGenerator{
		outputFile: outputFile,
		s:          s,
	}
}

func (eg *endpointsGenerator) generateCode() {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", eg.s.GetServiceName())
	fmt.Fprintf(&code, "import(\n%q\n%q\n%q\n%q\n%q\n)\n", "github.com/go-kit/kit/endpoint", "context", "crypto/md5", "encoding/json", "encoding/hex")

	for _, endpoint := range eg.s.Endpoints {

		fmt.Fprintf(&code, "\ntype %sRequest struct{\n", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s %s `json:%q`\n", eg.s.GetVariableName(arg, false), eg.s.GetType(arg), eg.s.GetVariableName(arg, true))

		}
		fmt.Fprintf(&code, "%s\n}", "")
		fmt.Fprintf(&code, "\ntype  %sResponse struct{\n", endpoint.GetName())
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s %s `json:%q`\n", eg.s.GetVariableName(out, false), eg.s.GetType(out), eg.s.GetVariableName(out, true))

		}
		fmt.Fprintf(&code, "\n}\n func make%sEndpoint(s %s)endpoint.Endpoint{\nreturn func(ctx context.Context, request interface{}) (interface{}, error) {\n", endpoint.GetName(), eg.s.GetInterfaceName())
		if endpoint.GetTransport()["method"] != "GET" {
			fmt.Fprintf(&code, "req := request.(%sRequest)\n", endpoint.GetName())
		}
		if endpoint.GetTransport()["method"] == "GET" && eg.s.RedisCache.Host != "" {
			fmt.Fprintf(&code, "req := request.(%sRequest)\n", endpoint.GetName())

			fmt.Fprintf(&code, "\nresp,error:=s.GetCache().Get(ctx,req.Hashcode())")
			fmt.Fprintf(&code, "\nif error!=nil{\n")
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "\n%s,", eg.s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "error:=s.%s(", endpoint.GetName())

			for _, arg := range endpoint.GetArgs() {
				fmt.Fprintf(&code, "req.%s,", eg.s.GetVariableName(arg, false))

			}

			fmt.Fprintf(&code, "ctx)\nresponse:= %sResponse{", endpoint.GetName())
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "%s: %s,", eg.s.GetVariableName(out, false), eg.s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "}")
			fmt.Fprintf(&code, "\ns.GetCache().Set(ctx,req.Hashcode(),response,%d)", endpoint.GetCacheTime())
			fmt.Fprintf(&code, "\nreturn response, error\n}")
			fmt.Fprintf(&code, "\nvar response %sResponse", endpoint.GetName())
			fmt.Fprintf(&code, "\n json.Unmarshal([]byte(resp),&response)")
			fmt.Fprintf(&code, "\nreturn response,nil")
			fmt.Fprintf(&code, "}}")

		} else {
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "\n%s,", eg.s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "error:=s.%s(", endpoint.GetName())

			for _, arg := range endpoint.GetArgs() {
				fmt.Fprintf(&code, "req.%s,", eg.s.GetVariableName(arg, false))

			}
			fmt.Fprintf(&code, "ctx)\nreturn %sResponse{", endpoint.GetName())
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "%s: %s,", eg.s.GetVariableName(out, false), eg.s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "}, error\n}\n}\n")

		}

	}
	fmt.Fprintf(&code, "%s", "\ntype Endpoints struct{\n ")
	for _, endpoint := range eg.s.Endpoints {

		fmt.Fprintf(&code, "%s endpoint.Endpoint\n", endpoint.GetName())

	}
	fmt.Fprintf(&code, "%s\n", "}")

	fmt.Fprintf(&code, "func MakeEndpoints(s %s)Endpoints{\n return Endpoints{\n", eg.s.GetInterfaceName())

	for _, endpoint := range eg.s.Endpoints {

		fmt.Fprintf(&code, "%s:make%sEndpoint(s),\n", endpoint.GetName(), endpoint.GetName())

	}

	fmt.Fprintf(&code, "%s\n}", "}")
	for _, endpoint := range eg.s.Endpoints {
		fmt.Fprintf(&code, "\nfunc (e %sRequest)Hashcode()string{", endpoint.GetName())
		fmt.Fprintf(&code, "\njson,_:=json.Marshal(e)")
		fmt.Fprintf(&code, "\nhasher:= md5.New()")
		fmt.Fprintf(&code, "\nhasher.Write(json)")
		fmt.Fprintf(&code, "\ncode:= hex.EncodeToString(hasher.Sum(nil))")
		fmt.Fprintf(&code, "\nreturn code")
		fmt.Fprintf(&code, "\n}")

	}

	eg.code = code.String()
}
func (eg *endpointsGenerator) generateFile(outputPath string) {
	var path string

	path = fmt.Sprintf("%s/%s.go", outputPath, eg.outputFile)

	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(eg.code)
	defer file.Close()
}
