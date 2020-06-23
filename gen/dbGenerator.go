package gen

import (
	"fmt"
	"strings"

	"github.com/samkhud/go-kit-code-generator/model"
)

func generateDb(s *model.Service) string {
	var code strings.Builder
	for _, m := range s.Models {
		fmt.Fprintf(&code, "CREATE TABLE %ss \n(", m.GetName(true))
		for _, attr := range m.Attributes {
			name := s.GetVariableName(m.Attr(attr), true)
			dbType := m.DBType(attr)
			if dbType == "" {
				dbType = "--TODO: fill type"
			}
			fmt.Fprintf(&code, "\n%s %s,", name, dbType)
		}
		fmt.Fprintf(&code, "\n);")

	}

	return code.String()

}
