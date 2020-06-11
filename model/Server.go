package model

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type Service struct {
	Name      string     `json:"name"`
	Endpoints []Endpoint `yaml:"endpoints"`
}
type Endpoint struct {
	Name      string    `yaml:"name"`
	Args      string    `yaml:"args"`
	Output    string    `yaml:"output"`
	Transport Transport `yaml:"transport"`
}
type Transport struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}

func (endpoint *Endpoint) GetArgs() []string {

	return strings.Split(endpoint.Args, ",")
}
func (endpoint *Endpoint) GetOutputs() []string {
	return strings.Split(endpoint.Output, ",")
}
func (endpoint *Endpoint) GetVariableName(in string, private bool) string {
	if private {
		return strcase.ToLowerCamel(strings.Split(strings.TrimSpace(in), " ")[0])
	} else {
		return strcase.ToCamel(strings.Split(strings.TrimSpace(in), " ")[0])
	}

}
func (endpoint *Endpoint) GetType(in string) string {
	return strings.Split(strings.TrimSpace(in), " ")[1]

}
func (endpoint *Endpoint) GetTransport() map[string]string {
	res := make(map[string]string)
	res["method"] = strings.ToUpper(strings.TrimSpace(endpoint.Transport.Method))
	res["path"] = strings.ToLower(strings.TrimSpace(endpoint.Transport.Path))
	return res
}
func (s *Endpoint) GetName() string {
	return strcase.ToCamel(s.Name)

}
func (s *Service) GetInterfaceName() string {
	return strcase.ToCamel(s.Name)

}
func (s *Service) GetServiceName() string {
	return strcase.ToLowerCamel(s.Name)

}
