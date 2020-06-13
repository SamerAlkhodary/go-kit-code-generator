package gen

import (
	"fmt"
	"go-kit-code-generator/model"
	"strings"
)

func mainCodeGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", s.GetServiceName())
	fmt.Fprintf(&code, "import(\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n)\n", "github.com/go-kit/kit/log", "github.com/go-kit/kit/log/level", "fmt", "flag", "net/http", "os", "os/signal", "syscall", "context", "database/sql")
	fmt.Fprintf(&code, "func Serve(){\nvar httpAddr= flag.String(%q,%q,%q)\nvar logger log.Logger\n{\n\n", "http", ":8080", "http listen address")
	fmt.Fprintf(&code, "\nlogger= log.NewLogfmtLogger(os.Stderr)\nlogger=log.NewSyncLogger(logger)\nlogger= log.With(logger,\n%q,%q,\n%q, log.DefaultTimestampUTC,\n%q, log.DefaultCaller,\n)\n}", "service", s.GetServiceName(), "time", "caller")
	fmt.Fprintf(&code, "\nlevel.Info(logger).Log(%q,%q)", "msg", "service started")
	fmt.Fprintf(&code, "\ndefer level.Info(logger).Log(%q,%q)", "msg", "service ended")
	if s.Repository.Value {
		fmt.Fprintf(&code, "\n var db *sql.DB")
		fmt.Fprintf(&code, "\n{\nvar err error")
		fmt.Fprintf(&code, "\ndb,err=sql.Open(%q,%q)", s.Repository.GetDB().GetName(), s.Repository.GetDB().GetAddress())
		fmt.Fprintf(&code, "\nif err!=nil{\nlevel.Error(logger).Log(%q,err)\nos.Exit(-1)\n}\n}\n", "exit")

		fmt.Fprintf(&code, "\n repository:= MakeNewRepository(db,logger)\n")

	}
	fmt.Fprintf(&code, "\nflag.Parse()\nctx:=context.Background()")
	if s.Repository.Value {
		fmt.Fprintf(&code, "\nvar service %s\n{\nservice= NewService(logger,repository)\n}\n", s.GetInterfaceName())

	} else {
		fmt.Fprintf(&code, "\nvar service %s\n{\nservice= NewService(logger)\n}\n", s.GetInterfaceName())

	}
	fmt.Fprintf(&code, "\nerrs:=make(chan error)\ngo func(){\n")
	fmt.Fprintf(&code, "\nc := make(chan os.Signal,1)\n signal.Notify(c,syscall.SIGINT, syscall.SIGTERM)\nerrs<- fmt.Errorf(%q,<-c)\n}()", "%s")
	fmt.Fprintf(&code, "\nendpoints:=MakeEndpoints(service)")
	fmt.Fprintf(&code, "\ngo func(){\nfmt.Println(%q,*httpAddr)", "Listening on port")
	fmt.Fprintf(&code, "\nhandler:=NewHTTPServer(ctx,endpoints)\n")
	fmt.Fprintf(&code, "\nerrs <- http.ListenAndServe(*httpAddr, handler)")
	fmt.Fprintf(&code, "\n}()")
	fmt.Fprintf(&code, "\nlevel.Error(logger).Log(%q, <-errs)\n}", "exit")

	return code.String()

}
