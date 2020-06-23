package gen

import (
	"fmt"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

func modelGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s", s.GetServiceName())
	for _, model := range s.Models {
		fmt.Fprintf(&code, "\ntype %s struct{\n", model.GetName(false))
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\n%s %s `json:%q`\n", s.GetVariableName(model.Attr(attr), false), s.GetType(attr), s.GetVariableName(model.Attr(attr), true))
		}
		fmt.Fprintf(&code, "}")
		fmt.Fprintf(&code, "\nfunc MakeNew%s(", model.GetName(false))
		for i, attr := range model.GetModelAttributes() {

			fmt.Fprintf(&code, "%s %s ", s.GetVariableName(model.Attr(attr), true), s.GetType(model.Attr(attr)))
			if i < len(model.Attr(attr))-3 {
				fmt.Fprintf(&code, ",")
			}
		}
		fmt.Fprintf(&code, ")*%s{\n return &%s{\n", model.GetName(false), model.GetName(false))
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\n%s : %s, \n", s.GetVariableName(model.Attr(attr), false), s.GetVariableName(model.Attr(attr), true))
		}
		fmt.Fprintf(&code, "}\n}")
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\nfunc(%s *%s)Get%s()%s{\nreturn %s.%s\n} \n", model.GetName(true), model.GetName(false), s.GetVariableName(model.Attr(attr), false), s.GetType(model.Attr(attr)), model.GetName(true), s.GetVariableName(model.Attr(attr), false))
			fmt.Fprintf(&code, "\nfunc(%s *%s)Set%s(new%s %s){\n%s.%s=new%s\n} \n", model.GetName(true), model.GetName(false), s.GetVariableName(model.Attr(attr), false), s.GetVariableName(model.Attr(attr), false), s.GetType(model.Attr(attr)), model.GetName(true), s.GetVariableName(model.Attr(attr), false), s.GetVariableName(model.Attr(attr), false))

		}

	}
	return code.String()
}
