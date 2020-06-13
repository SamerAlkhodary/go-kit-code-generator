package gen

import (
	"fmt"
	"go-kit-code-generator/model"
	"strings"
)

func endpointsGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", s.GetServiceName())
	fmt.Fprintf(&code, "import(\n%q\n%q\n)\n", "github.com/go-kit/kit/endpoint", "context")

	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "type %sRequest struct{\n", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s %s `json:%q`\n", s.GetVariableName(arg, false), s.GetType(arg), s.GetVariableName(arg, true))

		}
		fmt.Fprintf(&code, "%s\n}", "")
		fmt.Fprintf(&code, "\ntype  %sResponse struct{\n", endpoint.GetName())
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s %s `json:%q`\n", s.GetVariableName(out, false), s.GetType(out), s.GetVariableName(out, true))

		}
		fmt.Fprintf(&code, "\n}\n func make%sEndpoint(s %s)endpoint.Endpoint{\nreturn func(ctx context.Context, request interface{}) (interface{}, error) {\nreq := request.(%sRequest)\n", endpoint.GetName(), s.GetInterfaceName(), endpoint.GetName())
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", s.GetVariableName(out, true))
		}
		fmt.Fprintf(&code, "error:=s.%s(", endpoint.GetName())

		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "req.%s,", s.GetVariableName(arg, false))

		}
		fmt.Fprintf(&code, "ctx)\nreturn %sResponse{", endpoint.GetName())
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s: %s,", s.GetVariableName(out, false), s.GetVariableName(out, true))
		}
		fmt.Fprintf(&code, "}, error\n}\n}\n")
	}
	fmt.Fprintf(&code, "%s", "type Endpoints struct{\n ")
	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "%s endpoint.Endpoint\n", endpoint.GetName())

	}
	fmt.Fprintf(&code, "%s\n", "}")

	fmt.Fprintf(&code, "func MakeEndpoints(s %s)Endpoints{\n return Endpoints{\n", s.GetInterfaceName())

	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "%s:make%sEndpoint(s),\n", endpoint.GetName(), endpoint.GetName())

	}

	fmt.Fprintf(&code, "%s\n}", "}")

	return code.String()
}
