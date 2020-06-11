package gen

import (
	"fmt"
	"strings"
)

func transportGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", s.GetServiceName())
	fmt.Fprintf(&code, "import(httptransport %q\n%q\n%q\n%q)\n", "github.com/go-kit/kit/transport/http", "context", "github.com/gorilla/mux", "net/http")
	fmt.Fprintf(&code, "func NewHTTPServer(ctx context.Context,endPoints Endpoints)http.Handler{\n")
	fmt.Fprintf(&code, "r:=mux.NewRouter()\nr.Use(commonMiddleware)\n")
	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "r.Method(%q).Path(%q).Handler(httptransport.NewServer(\nendpoints.%s,\ndecode%sRequest,\nencode%sResponse,\n))\n", endpoint.GetTransport()["method"], endpoint.GetTransport()["path"], endpoint.GetName(), endpoint.GetName(), endpoint.GetName())

	}
	fmt.Fprintf(&code, "\nreturn r")

	fmt.Fprintf(&code, "\n}\n")
	fmt.Fprintf(&code, "func commonMiddleware(next http.Handler)http.Handler{\n")
	fmt.Fprintf(&code, " return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){\nw.Header().Add(%q,%q)\nnext.ServeHTTP(w,r)\n})}\n", "Content-Type", "application/json")

	return code.String()
}
