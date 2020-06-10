package model

import "strings"

type Service struct {
	Name      string     `json:"name"`
	Endpoints []Endpoint `yaml:"endpoints"`
}
type Endpoint struct {
	Name   string `yaml:"name"`
	Args   string `yaml:"args"`
	Output string `yaml:"output"`
}

func (endpoint *Endpoint) GetArgs() []string {
	return strings.Split(endpoint.Args, ",")
}
func (endpoint *Endpoint) GetOutputs() []string {
	return strings.Split(endpoint.Output, ",")
}
