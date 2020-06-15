package gen

import (
	"fmt"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

func repositroyGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s", s.GetServiceName())
	fmt.Fprintf(&code, "\nimport(\n%q\n%q\n%q\n)", "context", "github.com/go-kit/kit/log", "database/sql")
	fmt.Fprintf(&code, " \ntype Repository interface{\n")
	for _, endpoint := range s.Endpoints {
		fmt.Fprintf(&code, "\n%s(", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s %s,", s.GetVariableName(arg, true), s.GetType(arg))

		}
		fmt.Fprintf(&code, "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", s.GetType(out))

		}
		fmt.Fprintf(&code, "error)")

	}
	fmt.Fprintf(&code, "\n}\n")
	fmt.Fprintf(&code, "type repository struct{\n db *sql.DB\nlogger log.Logger\n}\n")
	fmt.Fprintf(&code, "\nfunc MakeNewRepository(db *sql.DB, logger log.Logger)Repository{\n ")
	fmt.Fprintf(&code, "\nreturn &repository{\ndb:db,\nlogger:log.With(logger,%q,%q ),\n}\n", "repository", "sql")
	fmt.Fprintf(&code, "\n}\n")
	for _, endpoint := range s.Endpoints {
		fmt.Fprintf(&code, "\nfunc(repo *repository)%s(", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s %s,", s.GetVariableName(arg, true), s.GetType(arg))
		}
		fmt.Fprintf(&code, "ctx context.Context)(")
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s,", s.GetType(out))
		}
		fmt.Fprintf(&code, "error){\n}\n")

	}

	return code.String()
}
