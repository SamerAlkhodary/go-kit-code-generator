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
	fmt.Fprintf(&code, "import(%q\n %q\n%q\n)\n", "context", "github.com/go-kit/kit/log", "github.com/go-kit/kit/log/level")

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
	if s.Repository {
		fmt.Fprintf(&code, "%s\n", "logger log.Logger\n repository Repository")

	} else {
		fmt.Fprintf(&code, "%s\n", "logger log.Logger\n")

	}

	fmt.Fprintf(&code, "%s\n", "}")
	if s.Repository {

		fmt.Fprintf(&code, "func NewService(logger log.Logger,repository Repository)%s{\n return &%s{\n logger:logger,\n repository:repository,\n}}\n", s.GetInterfaceName(), s.GetServiceName())
	} else {
		fmt.Fprintf(&code, "func NewService(logger log.Logger)%s{\n return &%s{\n logger:logger,\n}}\n", s.GetInterfaceName(), s.GetServiceName())

	}

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
		if s.Repository {
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "logger:= log.With(s.logger,%q,%q)\n", "method", endpoint.GetName())
				fmt.Fprintf(&code, "%s,", s.GetVariableName(out, true))

			}
			fmt.Fprintf(&code, "err:= s.repository.%s(", endpoint.GetName())
			for _, arg := range endpoint.GetArgs() {
				fmt.Fprintf(&code, "%s,", s.GetVariableName(arg, true))
			}
			fmt.Fprintf(&code, "ctx)\n")
			fmt.Fprintf(&code, "\nif err!=nil{")
			fmt.Fprintf(&code, "\nlevel.Error(logger).Log(%q,err)\n", "err")
			fmt.Fprintf(&code, "\nreturn nil,err \n")
			fmt.Fprintf(&code, "\n}\n")
			fmt.Fprintf(&code, "\nlogger.Log(%q)\n", endpoint.GetName())
			fmt.Fprintf(&code, "\nreturn ")
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "%s,", s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "nil")

		} else {
			fmt.Fprintf(&code, "logger:= log.With(s.logger,%q,%q)\n//TODO: implement\n", "method", endpoint.GetName())

		}

		fmt.Fprintf(&code, "}\n")

	}
	return code.String()
}
