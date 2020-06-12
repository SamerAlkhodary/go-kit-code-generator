package gen

import (
	"fmt"
	"go-kit-code-generator/model"
	"strings"
)

func serviceGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", s.GetServiceName())
	fmt.Fprintf(&code, "import(%q\n %q)\n", "context", "github.com/go-kit/kit/log")

	fmt.Fprintf(&code, "type %s interface{\n", s.GetInterfaceName())
	for _, endpoint := range s.Endpoints {
		fmt.Fprintf(&code, "%s(", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s,", arg)
		}
		fmt.Fprintf(&code, "%s", "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", s.GetType(out))
		}
		fmt.Fprintf(&code, "error)\n")

	}
	fmt.Fprintf(&code, "%s\n", "}")

	fmt.Fprintf(&code, "type %s struct{\n", s.GetServiceName())
	fmt.Fprintf(&code, "%s\n", "logger log.Logger")

	fmt.Fprintf(&code, "%s\n", "}")
	fmt.Fprintf(&code, "func NewService(logger log.Logger)%s{\n return &%s{\n logger:logger,\n}}\n", s.GetInterfaceName(), s.GetServiceName())
	for _, endpoint := range s.Endpoints {
		fmt.Fprintf(&code, "func(s *%s)%s(", s.GetServiceName(), endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s,", arg)
		}
		fmt.Fprintf(&code, "%s", "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", s.GetType(out))
		}
		fmt.Fprintf(&code, "error){\n")
		fmt.Fprintf(&code, "Logger:= log.With(s.logger,%q,%q)\n//TODO: implement\n", "method", endpoint.GetName())

		fmt.Fprintf(&code, "}\n")

	}
	return code.String()
}
