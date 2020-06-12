package gen

import (
	"fmt"
	"go-kit-code-generator/model"
	"strings"
)

func modelGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s", s.GetServiceName())
	for _, model := range s.Models {
		fmt.Fprintf(&code, "\ntype %s struct{\n", model.GetName(false))
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\n%s %s `json:%q`\n", s.GetVariableName(attr, true), s.GetType(attr), s.GetVariableName(attr, true))
		}
		fmt.Fprintf(&code, "}")
		fmt.Fprintf(&code, "\nfunc MakeNew%s(", model.GetName(false))
		for i, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "%s %s ", s.GetVariableName(attr, true), s.GetType(attr))
			if i < len(attr)-3 {
				fmt.Fprintf(&code, ",")
			}
		}
		fmt.Fprintf(&code, ")*%s{\n return &%s{\n", model.GetName(false), model.GetName(false))
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\n%s : %s, \n", s.GetVariableName(attr, true), s.GetVariableName(attr, true))
		}
		fmt.Fprintf(&code, "}\n}")
		for _, attr := range model.GetModelAttributes() {
			fmt.Fprintf(&code, "\nfunc(%s *%s)Get%s()%s{\nreturn %s.%s\n} \n", model.GetName(true), model.GetName(false), s.GetVariableName(attr, false), s.GetType(attr), model.GetName(true), s.GetVariableName(attr, true))
			fmt.Fprintf(&code, "\nfunc(%s *%s)Set%s(new%s %s){\n%s.%s=new%s\n} \n", model.GetName(true), model.GetName(false), s.GetVariableName(attr, false), s.GetVariableName(attr, false), s.GetType(attr), model.GetName(true), s.GetVariableName(attr, true), s.GetVariableName(attr, false))

		}

	}
	return code.String()
}
