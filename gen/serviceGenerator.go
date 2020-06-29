package gen

import (
	"fmt"
	"log"
	"os"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

type serviceGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createServiceGenerator(s model.Service, outputFile string) fileGenerator {
	return &serviceGenerator{
		outputFile: outputFile,
		s:          s,
	}
}

func (sg *serviceGenerator) run(outputPath string) {
	sg.generateCode()
	sg.generateFile(outputPath)
}

func (sg *serviceGenerator) generateCode() {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", sg.s.GetServiceName())
	fmt.Fprintf(&code, "import(%q\n %q\n%q\n)\n", "context", "github.com/go-kit/kit/log", "github.com/go-kit/kit/log/level")

	fmt.Fprintf(&code, "type %s interface{\n", sg.s.GetInterfaceName())
	for _, endpoint := range sg.s.Endpoints {
		fmt.Fprintf(&code, "\n%s(", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s,", arg)
		}
		fmt.Fprintf(&code, "%s", "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", sg.s.GetType(out))
		}
		fmt.Fprintf(&code, "error)")

	}
	if sg.s.RedisCache.GetHost() != "" {
		fmt.Fprintf(&code, "\nGetCache()Cache")

	}

	fmt.Fprintf(&code, "%s\n", "}")

	fmt.Fprintf(&code, "type %s struct{\n", sg.s.GetServiceName())
	if sg.s.Repository.Value {
		if sg.s.RedisCache.GetHost() != "" {
			fmt.Fprintf(&code, "%s\n", "logger log.Logger\n repository Repository\ncache Cache")

		} else {
			fmt.Fprintf(&code, "%s\n", "logger log.Logger\n repository Repository")

		}

	} else {
		if sg.s.RedisCache.GetHost() != "" {
			fmt.Fprintf(&code, "%s\n", "logger log.Logger\ncache Cache")

		} else {
			fmt.Fprintf(&code, "%s\n", "logger log.Logger\n")

		}

	}

	fmt.Fprintf(&code, "%s\n", "}")
	if sg.s.Repository.Value {
		if sg.s.RedisCache.GetHost() != "" {
			fmt.Fprintf(&code, "func NewService(logger log.Logger,repository Repository,cache Cache)%s{\n return &%s{\n logger:logger,\n repository:repository,\ncache:cache,\n}}\n", sg.s.GetInterfaceName(), sg.s.GetServiceName())

		} else {
			fmt.Fprintf(&code, "func NewService(logger log.Logger,repository Repository)%s{\n return &%s{\n logger:logger,\n repository:repository,\n}}\n", sg.s.GetInterfaceName(), sg.s.GetServiceName())
		}

	} else {
		if sg.s.RedisCache.GetHost() != "" {
			fmt.Fprintf(&code, "func NewService(logger log.Logger)%s{\n return &%s{\n logger:logger,\ncache:cache,\n}}\n", sg.s.GetInterfaceName(), sg.s.GetServiceName())
		} else {
			fmt.Fprintf(&code, "func NewService(logger log.Logger)%s{\n return &%s{\n logger:logger,\n}}\n", sg.s.GetInterfaceName(), sg.s.GetServiceName())
		}

	}

	for _, endpoint := range sg.s.Endpoints {
		fmt.Fprintf(&code, "\nfunc(s *%s)%s(", sg.s.GetServiceName(), endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s,", arg)
		}
		fmt.Fprintf(&code, "%s", "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", sg.s.GetType(out))
		}
		fmt.Fprintf(&code, "error){\n")
		if sg.s.Repository.Value {
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "logger:= log.With(s.logger,%q,%q)\n", "method", endpoint.GetName())
				fmt.Fprintf(&code, "%s,", sg.s.GetVariableName(out, true))

			}
			fmt.Fprintf(&code, "err:= s.repository.%s(", endpoint.GetName())
			for _, arg := range endpoint.GetArgs() {
				fmt.Fprintf(&code, "%s,", sg.s.GetVariableName(arg, true))
			}
			fmt.Fprintf(&code, "ctx)\n")
			fmt.Fprintf(&code, "\nif err!=nil{")
			fmt.Fprintf(&code, "\nlevel.Error(logger).Log(%q,err)\n", "err")
			fmt.Fprintf(&code, "\n//TODO: fix return")
			fmt.Fprintf(&code, "\nreturn nil,err \n")
			fmt.Fprintf(&code, "\n}\n")
			fmt.Fprintf(&code, "\nlogger.Log(%q)\n", endpoint.GetName())

			fmt.Fprintf(&code, "\nreturn ")
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "%s,", sg.s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "nil")

		} else {
			fmt.Fprintf(&code, "logger:= log.With(s.logger,%q,%q)\n//TODO: implement\n", "method", endpoint.GetName())

		}

		fmt.Fprintf(&code, "}")

	}
	if sg.s.RedisCache.Host != "" {
		fmt.Fprintf(&code, "\nfunc(s *%s)GetCache()Cache{\n", sg.s.GetServiceName())
		fmt.Fprintf(&code, "\nreturn s.cache\n}")

	}
	sg.code = code.String()

}
func (sg *serviceGenerator) generateFile(outputPath string) {
	var path string

	path = fmt.Sprintf("%s/%s.go", outputPath, sg.outputFile)

	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(sg.code)
	defer file.Close()

}
