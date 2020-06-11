package gen

import (
	"fmt"
	"go-kit-code-generator/model"
	"strings"
)

func mainCodeGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package main\n")
	fmt.Fprintf(&code, "import(\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n)\n", "github.com/go-kit/kit/log", "github.com/go-kit/kit/log/level", "fmt", "flag", "net/http", "os", "os/signal", "syscall", "context")
	fmt.Fprintf(&code, "func main(){\nvar logger log.Logger\n{\n\n")
	fmt.Fprintf(&code, "logger= log.NewLogfmtLogger(os.Stderr)\nlogger=log.NewSyncLogger(logger)\nlogger= log.With(logger,\n%q,%q,\n%q, log.DefaultTimestampUTC,\n%q, log.DefaultCaller,\n)\n}", "service", s.GetServiceName(), "time", "caller")
	fmt.Fprintf(&code, "\nlevel.Info(logger).Log(%q,%q)", "msg", "service started")
	fmt.Fprintf(&code, "\ndefer level.Info(logger).Log(%q,%q)", "msg", "service ended")
	fmt.Fprintf(&code, "\nflag.Pars()\nctx:=context.Background()")
	fmt.Fprintf(&code, "\nvar service %s.%s\n{\nservice= %s.NewServer(logger)\n}\n", s.GetServiceName(), s.GetInterfaceName(), s.GetServiceName())

	return code.String()

}
