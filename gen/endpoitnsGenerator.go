package gen

import (
	"fmt"
	"strings"

	"github.com/samkhud/go-kit-code-generator/model"
)

func endpointsGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "package %s\n", s.GetServiceName())
	fmt.Fprintf(&code, "import(\n%q\n%q\n%q\n%q\n%q\n)\n", "github.com/go-kit/kit/endpoint", "context", "crypto/md5", "encoding/json", "encoding/hex")

	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "\ntype %sRequest struct{\n", endpoint.GetName())
		for _, arg := range endpoint.GetArgs() {
			fmt.Fprintf(&code, "%s %s `json:%q`\n", s.GetVariableName(arg, false), s.GetType(arg), s.GetVariableName(arg, true))

		}
		fmt.Fprintf(&code, "%s\n}", "")
		fmt.Fprintf(&code, "\ntype  %sResponse struct{\n", endpoint.GetName())
		for _, out := range endpoint.GetOutputs() {
			fmt.Fprintf(&code, "%s %s `json:%q`\n", s.GetVariableName(out, false), s.GetType(out), s.GetVariableName(out, true))

		}
		fmt.Fprintf(&code, "\n}\n func make%sEndpoint(s %s)endpoint.Endpoint{\nreturn func(ctx context.Context, request interface{}) (interface{}, error) {\nreq := request.(%sRequest)\n", endpoint.GetName(), s.GetInterfaceName(), endpoint.GetName())
		if endpoint.GetTransport()["method"] == "GET" && s.RedisCache.Host != "" {

			fmt.Fprintf(&code, "\nresp,error:=s.GetCache().Get(ctx,req.Hashcode())")
			fmt.Fprintf(&code, "\nif error!=nil{\n")
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "\n%s,", s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "error:=s.%s(", endpoint.GetName())

			for _, arg := range endpoint.GetArgs() {
				fmt.Fprintf(&code, "req.%s,", s.GetVariableName(arg, false))

			}

			fmt.Fprintf(&code, "ctx)\nresponse:= %sResponse{", endpoint.GetName())
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "%s: %s,", s.GetVariableName(out, false), s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "}")
			fmt.Fprintf(&code, "\ns.GetCache().Set(req.Hashcode(),response)")
			fmt.Fprintf(&code, "\nreturn response, error\n}")
			fmt.Fprintf(&code, "\nvar response %sResponse", endpoint.GetName())
			fmt.Fprintf(&code, "\n json.Unmarshal(resp,&response)")
			fmt.Fprintf(&code, "\nreturn response,nil")
			fmt.Fprintf(&code, "}}")

		} else {
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "\n%s,", s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "error:=s.%s(", endpoint.GetName())

			for _, arg := range endpoint.GetArgs() {
				fmt.Fprintf(&code, "req.%s,", s.GetVariableName(arg, false))

			}
			fmt.Fprintf(&code, "ctx)\nreturn %sResponse{", endpoint.GetName())
			for _, out := range endpoint.GetOutputs() {
				fmt.Fprintf(&code, "%s: %s,", s.GetVariableName(out, false), s.GetVariableName(out, true))
			}
			fmt.Fprintf(&code, "}, error\n}\n}\n")

		}

	}
	fmt.Fprintf(&code, "%s", "\ntype Endpoints struct{\n ")
	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "%s endpoint.Endpoint\n", endpoint.GetName())

	}
	fmt.Fprintf(&code, "%s\n", "}")

	fmt.Fprintf(&code, "func MakeEndpoints(s %s)Endpoints{\n return Endpoints{\n", s.GetInterfaceName())

	for _, endpoint := range s.Endpoints {

		fmt.Fprintf(&code, "%s:make%sEndpoint(s),\n", endpoint.GetName(), endpoint.GetName())

	}

	fmt.Fprintf(&code, "%s\n}", "}")
	for _, endpoint := range s.Endpoints {
		fmt.Fprintf(&code, "\nfunc (e %sRequest)Hashcode()string{", endpoint.GetName())
		fmt.Fprintf(&code, "\njson,_:=json.Marshal(e)")
		fmt.Fprintf(&code, "\nhasher:= md5.New()")
		fmt.Fprintf(&code, "\nhasher.Write(json)")
		fmt.Fprintf(&code, "\ncode:= hex.EncodeToString(hasher.Sum(nil))")
		fmt.Fprintf(&code, "\nreturn code")
		fmt.Fprintf(&code, "\n}")

	}

	return code.String()
}
