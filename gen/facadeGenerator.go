package gen

import (
	"fmt"
	"log"
	"os"

	"github.com/samkhud/go-kit-code-generator/model"

	"strings"
)

type mainCodeGenerator struct {
	outputFile string
	s          model.Service
	code       string
}

func createMainGenerator(s model.Service, outputFile string) fileGenerator {
	return &mainCodeGenerator{
		outputFile: outputFile,
		s:          s,
	}
}
func (mg *mainCodeGenerator) GetFileName() string {
	return mg.outputFile
}

func (mg *mainCodeGenerator) run(outputPath string) {
	mg.generateCode()
	mg.generateFile(outputPath)
}
func (mg *mainCodeGenerator) generateCode() {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", mg.s.GetServiceName())

	fmt.Fprintf(&code, "import(\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n%q\n\n", "github.com/go-kit/kit/log", "github.com/go-kit/kit/log/level", "fmt", "flag", "net/http", "os", "os/signal", "syscall", "context", "database/sql")
	if mg.s.HasKeys() {
		fmt.Fprintf(&code, "\n%q", "crypto/tls")
	}
	fmt.Fprintf(&code, "\n)")

	fmt.Fprintf(&code, "\nfunc Serve(){\nvar httpAddr= flag.String(%q,%q,%q)\nvar logger log.Logger\n{\n\n", "http", ":8080", "http listen address")
	fmt.Fprintf(&code, "\nlogger= log.NewLogfmtLogger(os.Stderr)\nlogger=log.NewSyncLogger(logger)\nlogger= log.With(logger,\n%q,%q,\n%q, log.DefaultTimestampUTC,\n%q, log.DefaultCaller,\n)\n}", "service", mg.s.GetServiceName(), "time", "caller")
	fmt.Fprintf(&code, "\nlevel.Info(logger).Log(%q,%q)", "msg", "service started")
	fmt.Fprintf(&code, "\ndefer level.Info(logger).Log(%q,%q)", "msg", "service ended")
	if mg.s.Repository.Value {
		fmt.Fprintf(&code, "\n var db *sql.DB")
		fmt.Fprintf(&code, "\n{\nvar err error")
		fmt.Fprintf(&code, "\ndb,err=sql.Open(%q,%q)", mg.s.Repository.GetDB().GetName(), mg.s.Repository.GetDB().GetAddress())
		fmt.Fprintf(&code, "\nif err!=nil{\nlevel.Error(logger).Log(%q,err)\nos.Exit(-1)\n}\n}\n", "exit")

		fmt.Fprintf(&code, "\n repository:= MakeNewRepository(db,logger)\n")
		if mg.s.RedisCache.GetHost() != "" {
			fmt.Fprintf(&code, "\ncache:= MakeNewRedisCache(%q,%q,%d)\n", mg.s.RedisCache.GetHost(), mg.s.RedisCache.Password, mg.s.RedisCache.Db)

		}

	}
	fmt.Fprintf(&code, "\nflag.Parse()\nctx:=context.Background()")
	if mg.s.Repository.Value {
		if mg.s.RedisCache.GetHost() != "" {
			fmt.Fprintf(&code, "\nvar service %s\n{\nservice= NewService(logger,repository,cache)\n}\n", mg.s.GetInterfaceName())

		} else {
			fmt.Fprintf(&code, "\nvar service %s\n{\nservice= NewService(logger,repository)\n}\n", mg.s.GetInterfaceName())

		}

	} else {
		if mg.s.RedisCache.GetHost() != "" {
			fmt.Fprintf(&code, "\nvar service %s\n{\nservice= NewService(logger,cache)\n}\n", mg.s.GetInterfaceName())

		} else {
			fmt.Fprintf(&code, "\nvar service %s\n{\nservice= NewService(logger)\n}\n", mg.s.GetInterfaceName())
		}
	}
	fmt.Fprintf(&code, "\nerrs:=make(chan error)\ngo func(){\n")
	fmt.Fprintf(&code, "\nc := make(chan os.Signal,1)\n signal.Notify(c,syscall.SIGINT, syscall.SIGTERM)\nerrs<- fmt.Errorf(%q,<-c)\n}()", "%s")
	fmt.Fprintf(&code, "\nendpoints:=MakeEndpoints(service)")
	fmt.Fprintf(&code, "\ngo func(){\nfmt.Println(%q,*httpAddr)", "Listening on port")
	if mg.s.HasKeys() {
		fmt.Fprintf(&code, "\n//TODO: fill cert and key names")
		fmt.Fprintf(&code, "\ncert,error:=tls.LoadX509KeyPair(%q,%q)", "../keys/*.crt", "../keys/*.key")
		fmt.Fprintf(&code, "\nerrs<-error")

	}

	fmt.Fprintf(&code, "\nhandler:=NewHTTPServer(ctx,endpoints)\n")
	fmt.Fprintf(&code, "\nserver:= &http.Server{\n")
	fmt.Fprintf(&code, "\nAddr:*httpAddr,")
	fmt.Fprintf(&code, "\nHandler: handler,")
	if mg.s.HasKeys() {
		fmt.Fprintf(&code, "\nTLSConfig: &tls.Config{")
		fmt.Fprintf(&code, "\nCertificates:[] tls.Certificate{cert},},")
	}
	fmt.Fprintf(&code, "\n}")
	if mg.s.HasKeys() {
		fmt.Fprintf(&code, "\nerrs <- server.ListenAndServeTLS(%q,%q)", "", "")

	} else {
		fmt.Fprintf(&code, "\nerrs <- server.ListenAndServe()")

	}

	fmt.Fprintf(&code, "\n}()")
	fmt.Fprintf(&code, "\nlevel.Error(logger).Log(%q, <-errs)\n}", "exit")
	mg.code = code.String()

}
func (mg *mainCodeGenerator) generateFile(outputPath string) {
	var path string

	path = fmt.Sprintf("%s/%s.go", outputPath, mg.outputFile)

	file, err := os.Create(path)
	if err != nil {
		log.Printf("error while creating file:%v", err)
	}

	file.WriteString(mg.code)
	defer file.Close()
}
