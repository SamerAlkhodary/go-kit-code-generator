package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/samkhud/go-kit-code-generator/model"
	"gopkg.in/yaml.v2"
)

type Parser struct {
}

func (parser *Parser) Parse(path string) (*model.Service, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading file (%s): %v", path, err)

	}
	service := model.Service{}
	er := yaml.Unmarshal([]byte(file), &service)
	if er != nil {
		return nil, fmt.Errorf("error in Yaml: %v", er)

	}

	return &service, nil

}
