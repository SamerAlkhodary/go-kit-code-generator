package gen

import (
	"fmt"
	"log"
	"os"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

type transportGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createTransportGenerator(s model.Service, outputFile string) fileGenerator {
	return &transportGenerator{
		outputFile: outputFile,
		s:          s,
	}
}

func (tg *transportGenerator) run(outputPath string) {

	tg.generateCode()
	tg.generateFile(outputPath)
}

func (tg *transportGenerator) generateCode() {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", tg.s.GetServiceName())
	fmt.Fprintf(&code, "import(httptransport %q\n%q\n%q\n%q)\n", "github.com/go-kit/kit/transport/http", "context", "github.com/gorilla/mux", "net/http")
	fmt.Fprintf(&code, "func NewHTTPServer(ctx context.Context,endpoints Endpoints)http.Handler{\n")
	fmt.Fprintf(&code, "r:=mux.NewRouter()\nr.Use(commonMiddleware)\n")
	for _, endpoint := range tg.s.Endpoints {

		fmt.Fprintf(&code, "r.Methods(%q).Path(%q).Handler(httptransport.NewServer(\nendpoints.%s,\ndecode%sRequest,\nencodeResponse,\n))\n", endpoint.GetTransport()["method"], endpoint.GetTransport()["path"], endpoint.GetName(), endpoint.GetName())

	}
	fmt.Fprintf(&code, "\nreturn r")

	fmt.Fprintf(&code, "\n}\n")
	fmt.Fprintf(&code, "func commonMiddleware(next http.Handler)http.Handler{\n")
	fmt.Fprintf(&code, " return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){\nw.Header().Add(%q,%q)\nnext.ServeHTTP(w,r)\n})}\n", "Content-Type", "application/json")

	tg.code = code.String()
}
func (tg *transportGenerator) generateFile(outputPath string) {
	var path string

	path = fmt.Sprintf("%s/%s.go", outputPath, tg.outputFile)

	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(tg.code)
	defer file.Close()

}
