package gen

import (
	"fmt"
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
	fmt.Fprintf(&code, "\nerrs:=make(chan error\ngo func(){\n")
	fmt.Fprintf(&code, "\nc := make(chan os.Signal,1)\n signal.Notify(c,syscall.SIGINT, syscall.SIGTERM\nerrs<- fmt.Errorf(%q,<-c)\n}\n()","%s")
	fmt.Fprintf(&code, "\ngo func(){\nfmt.Println(%q,*httpAddr)","Listening on port")
	fmt.Fprintf(&code, "\nhandler:=%s.NewHTTPServer(ctx,endpoints)\n")
	fmt.Fprintf(&code, "\n}()")
	fmt.Fprintf(&code, "\nlevel.Error(looger).Log(%q, <-errs)\n}\n}","exit")





	
	return code.String()

}
