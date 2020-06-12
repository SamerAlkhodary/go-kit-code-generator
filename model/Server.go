package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type Service struct {
	Name       string     `yaml:"name"`
	Endpoints  []Endpoint `yaml:"endpoints"`
	Models     []Model    `yaml:"model"`
	Repository bool       `yaml:"repository"`
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
type Model struct {
	Name       string `yaml:"name"`
	Attributes string `yaml:"attr"`
}

var goTypes = make(map[string]bool)

func init() {
	goTypes["string"] = true
	goTypes["int"] = true
	goTypes["float32"] = true
	goTypes["float64"] = true
	goTypes["byte"] = true
	goTypes["uint"] = true
	goTypes["uint8"] = true
	goTypes["uint16"] = true
	goTypes["uint32"] = true
	goTypes["uint64"] = true
	goTypes["int8"] = true
	goTypes["int16"] = true
	goTypes["int32"] = true
	goTypes["int64"] = true
	goTypes["bool"] = true

}

var compileErr = errors.New("Compiling error")

func (endpoint *Endpoint) GetArgs() []string {

	return strings.Split(strings.TrimSpace(endpoint.Args), ",")
}
func (endpoint *Endpoint) GetOutputs() []string {
	return strings.Split(endpoint.Output, ",")
}
func (s *Service) GetVariableName(in string, private bool) string {
	if private {
		return strcase.ToLowerCamel(strings.Split(strings.TrimSpace(in), " ")[0])
	} else {
		return strcase.ToCamel(strings.Split(strings.TrimSpace(in), " ")[0])
	}

}
func (m *Model) GetModelAttributes() []string {
	return strings.Split(strings.TrimSpace(m.Attributes), ",")
}
func (m *Model) GetName(private bool) string {
	if private {
		return strcase.ToLowerCamel(strings.TrimSpace(m.Name))

	}
	return strcase.ToCamel(strings.TrimSpace(m.Name))

}
func (s *Service) GetType(in string) string {
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
func (s *Service) CheckForError() error {
	var err error
	err = checkServiceError(s)
	err = checkEndpointError(s)
	err = checkModelError(s)

	return err
}

func checkServiceError(s *Service) error {
	if s.Name == "" {
		return fmt.Errorf("Missing service name:%v", compileErr)
	}
	if len(s.Endpoints) == 0 {
		return fmt.Errorf("Missing endpoints:%v", compileErr)
	}

	return nil

}
func (s *Service) Apply() {
	for _, m := range s.Models {
		goTypes[m.GetName(false)] = true
	}
}

func checkEndpointError(s *Service) error {
	for _, endpoint := range s.Endpoints {
		if endpoint.GetName() == "" {
			return fmt.Errorf("Missing endpoint name:%v", compileErr)

		}
		for _, arg := range endpoint.GetArgs() {
			if len(strings.Split(strings.TrimSpace(arg), " ")) < 2 {
				return fmt.Errorf("Mising type or variable name in %s endpoint  :%v", endpoint.GetName(), compileErr)

			}
			if goTypes[s.GetType(arg)] == false {
				return fmt.Errorf("Unrecognised type %q in %s endpoint: %v", s.GetType(arg), endpoint.GetName(), compileErr)
			}
			if endpoint.GetTransport()["path"] == "" || endpoint.GetTransport()["method"] == "" {
				return fmt.Errorf("Missing transport info in %s endpoint: %v", endpoint.GetName(), compileErr)
			}
		}
		for _, out := range endpoint.GetOutputs() {
			if len(strings.Split(strings.TrimSpace(out), " ")) < 2 {
				return fmt.Errorf("Mising type or variable name in %s endpoint  :%v", endpoint.GetName(), compileErr)

			}
			if goTypes[s.GetType(out)] == false {
				return fmt.Errorf("Unrecognised type %q in %s endpoint: %v", s.GetType(out), endpoint.GetName(), compileErr)
			}

		}

	}
	return nil

}
func checkModelError(s *Service) error {
	if len(s.Models) == 0 {
		return nil
	}
	for _, m := range s.Models {
		if m.GetName(false) == "" {
			return fmt.Errorf("Missing model name:%v", compileErr)

		}
		for _, attr := range m.GetModelAttributes() {
			if len(strings.Split(strings.TrimSpace(attr), " ")) < 2 {
				return fmt.Errorf("Mising type or variable name in %s endpoint  :%v", m.GetName(false), compileErr)

			}
			if goTypes[s.GetType(attr)] == false {
				return fmt.Errorf("Unrecognised type %q in %s endpoint: %v", s.GetType(attr), m.GetName(false), compileErr)
			}

		}

	}
	return nil

}
