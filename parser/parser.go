package parser

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"services/generator/model"
)

type Parser struct {
}

func (parser *Parser) Parse(path string) *model.Service {
	//TODO: Parses the file and returns a service pointer
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error while reading file: %v", err)

	}
	service := model.Service{}
	yaml.Unmarshal([]byte(file), &service)
	fmt.Println(service.Name)
	fmt.Println(service.Endpoints)
	return &service

}
