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
	Models     []*Model   `yaml:"model"`
	Repository Repository `yaml:"repository"`
	RedisCache Cache      `yaml:"redis_cache"`
}
type Cache struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}
type Repository struct {
	Value bool `yaml:"value"`
	DB    DB   `yaml:"db"`
}
type DB struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
}
type Endpoint struct {
	Name      string    `yaml:"name"`
	Args      []string  `yaml:"args"`
	Output    []string  `yaml:"output"`
	CacheTime int       `yaml:"cache_time"`
	Transport Transport `yaml:"transport"`
}
type Transport struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}
type Model struct {
	Name             string   `yaml:"name"`
	StringAttributes []string `yaml:"attr"`
	Attributes       []Attribute
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
	goTypes["string"] = true
	goTypes["int"] = true
	goTypes["[]float32"] = true
	goTypes["[]float64"] = true
	goTypes["[]byte"] = true
	goTypes["[]uint"] = true
	goTypes["[]uint8"] = true
	goTypes["[]uint16"] = true
	goTypes["[]uint32"] = true
	goTypes["[]uint64"] = true
	goTypes["[]int8"] = true
	goTypes["[]int16"] = true
	goTypes["[]int32"] = true
	goTypes["[]int64"] = true
	goTypes["[]bool"] = true
	goTypes["[]string"] = true

}
func (s *Service) IsNativeType(typ string) bool {

	return goTypes[typ]

}
func (s *Service) IsAddedType(typ string) bool {
	if val, ok := goTypes[typ]; ok {
		return !val
	}
	return false

}
func (s *Service) IsArray(typ string) bool {
	if _, ok := goTypes[typ]; ok {
		return strings.Contains(typ, "[]")
	}
	return false

}
func (c *Cache) GetHost() string {
	return c.Host
}
func (r *Repository) GetDB() DB {
	return r.DB
}
func (d DB) GetName() string {
	return strings.ToLower(d.Name)
}
func (d DB) GetAddress() string {
	return d.Address
}

var compileErr = errors.New("Compiling error")

func (endpoint *Endpoint) GetArgs() []string {

	return filterEmpty(endpoint.Args)
}
func (endpoint *Endpoint) GetCacheTime() int {
	return endpoint.CacheTime
}
func (endpoint *Endpoint) GetOutputs() []string {
	return filterEmpty(endpoint.Output)
}
func (s *Service) GetVariableName(in string, private bool) string {
	if private {
		return strcase.ToLowerCamel(strings.Split(strings.TrimSpace(in), " ")[0])
	} else {
		return strcase.ToCamel(strings.Split(strings.TrimSpace(in), " ")[0])
	}

}
func (m *Model) GetModelAttributes() []Attribute {
	return m.Attributes
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
	if err != nil {
		return err
	}
	err = checkEndpointError(s)
	if err != nil {
		return err
	}
	err = checkModelError(s)
	if err != nil {
		return err
	}

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
		goTypes[m.GetName(false)] = false
		str := fmt.Sprintf("[]%s", m.GetName(false))
		goTypes[str] = false
		for _, attr := range m.StringAttributes {
			data := strings.Split(attr, "db=")
			if len(data) > 1 {
				m.Attributes = append(m.Attributes, *NewAttribute(s.GetVariableName(data[0], true), s.GetType(data[0]), data[1]))
			} else {
				m.Attributes = append(m.Attributes, *NewAttribute(s.GetVariableName(data[0], true), s.GetType(data[0]), "--TODO: write type"))
			}

		}

	}
}

func checkEndpointError(s *Service) error {
	for _, endpoint := range s.Endpoints {
		if endpoint.GetName() == "" {
			return fmt.Errorf("Missing endpoint name:%v", compileErr)

		}
		for _, arg := range endpoint.GetArgs() {
			if strings.TrimSpace(arg) != "" {
				if len(strings.Split(strings.TrimSpace(arg), " ")) < 2 {

					return fmt.Errorf("Mising type or variable name in %s endpoint  :%v", endpoint.GetName(), compileErr)
				}

			}

			if !s.IsNativeType(s.GetType(arg)) && !s.IsAddedType(s.GetType(arg)) {

				return fmt.Errorf("Unrecognised type %q in %s endpoint: %v", s.GetType(arg), endpoint.GetName(), compileErr)
			}
			if endpoint.GetTransport()["path"] == "" || endpoint.GetTransport()["method"] == "" {
				return fmt.Errorf("Missing transport info in %s endpoint: %v", endpoint.GetName(), compileErr)
			}
		}
		for _, out := range endpoint.GetOutputs() {
			if strings.TrimSpace(out) != "" {
				if len(strings.Split(strings.TrimSpace(out), " ")) < 2 {
					return fmt.Errorf("Mising type or variable name in %s endpoint  :%v", endpoint.GetName(), compileErr)

				}
			}
			if !s.IsNativeType(s.GetType(out)) && !s.IsAddedType(s.GetType(out)) {
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
		for _, attr := range m.StringAttributes {
			if len(strings.Split(strings.TrimSpace(attr), " ")) < 2 {
				return fmt.Errorf("Mising type or variable name in %s model  :%v", m.GetName(false), compileErr)

			}
			if !s.IsNativeType(s.GetType(attr)) && !s.IsAddedType(s.GetType(attr)) {
				return fmt.Errorf("Unrecognised type %q in %s model: %v", s.GetType(attr), m.GetName(false), compileErr)
			}
		}
	}
	return nil

}
func filterEmpty(arr []string) []string {
	tmp := arr[:0]
	for _, elem := range arr {
		if elem != "" {
			tmp = append(tmp, elem)

		}
	}
	return tmp
}
