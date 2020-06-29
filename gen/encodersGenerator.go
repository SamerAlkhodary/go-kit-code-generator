package gen

import (
	"fmt"
	"log"
	"os"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

type encoderDecoderGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createEncodersGenerator(s model.Service, outputFile string) fileGenerator {
	return &encoderDecoderGenerator{
		outputFile: outputFile,
		s:          s,
	}
}

func (eg *encoderDecoderGenerator) run(outputPath string) {
	eg.generateCode()
	eg.generateFile(outputPath)
}

func (eg *encoderDecoderGenerator) generateCode() {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", eg.s.GetServiceName())
	fmt.Fprintf(&code, "import(\n%q\n%q\n %q\n)\n", "encoding/json", "context", "net/http")
	fmt.Fprintf(&code, "func encodeResponse(ctx context.Context, w http.ResponseWriter,response interface{})error{\n")
	fmt.Fprintf(&code, "return json.NewEncoder(w).Encode(response)\n}")
	for _, endpoint := range eg.s.Endpoints {
		if endpoint.GetTransport()["method"] != "GET" {
			fmt.Fprintf(&code, "\nfunc decode%sRequest(ctx context.Context, r *http.Request)(interface{},error){\n", endpoint.GetName())
			fmt.Fprintf(&code, "var request %sRequest\nerr:=json.NewDecoder(r.Body).Decode(&request)\n", endpoint.GetName())
			fmt.Fprintf(&code, "if err!=nil{\n return nil,err\n}\n return request,nil\n}\n")
		} else {
			fmt.Fprintf(&code, "\nfunc decode%sRequest(ctx context.Context, r *http.Request)(interface{},error){\n", endpoint.GetName())
			if len(endpoint.GetArgs()) != 0 {
				fmt.Fprintf(&code, "\n vars:= mux.Vars(r)")

			}
			fmt.Fprintf(&code, "\nvar request %sRequest\n request= %sRequest{\n ", endpoint.GetName(), endpoint.GetName())

			for _, arg := range endpoint.GetArgs() {
				fmt.Fprintf(&code, "%s: vars[%q]", eg.s.GetVariableName(arg, false), eg.s.GetVariableName(arg, true))
			}
			fmt.Fprintf(&code, "}\nreturn request,nil\n}\n")
		}
	}

	eg.code = code.String()
}

func (eg *encoderDecoderGenerator) generateFile(outputPath string) {
	var path string

	path = fmt.Sprintf("%s/%s.go", outputPath, eg.outputFile)

	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(eg.code)
	defer file.Close()

}
