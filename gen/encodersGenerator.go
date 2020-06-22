package gen

import (
	"fmt"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

func encoderDecoderGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", s.GetServiceName())
	fmt.Fprintf(&code, "import(\n%q\n%q\n %q\n)\n", "encoding/json", "context", "net/http")
	fmt.Fprintf(&code, "func encodeResponse(ctx context.Context, w http.ResponseWriter,response interface{})error{\n")
	fmt.Fprintf(&code, "return json.NewEncoder(w).Encode(response)\n}")
	for _, endpoint := range s.Endpoints {
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
				fmt.Fprintf(&code, "%s: vars[%q]", s.GetVariableName(arg, false), s.GetVariableName(arg, true))
			}
			fmt.Fprintf(&code, "}\nreturn request,nil\n}\n")
		}
	}

	return code.String()
}
