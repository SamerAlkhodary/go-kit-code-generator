package model

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type Attribute struct {
	Name     string
	DataType string
	DBType   string
}

func NewAttribute(name string, dataType string, dbType string) *Attribute {
	return &Attribute{
		Name:     strings.TrimSpace(name),
		DataType: strings.TrimSpace(dataType),
		DBType:   strings.TrimSpace(dbType),
	}
}
func (a *Attribute) GetName(private bool) string {
	if private {
		return a.Name
	}
	return strcase.ToCamel(a.Name)

}
func (a *Attribute) GetType() string {
	return a.DataType

}
func (a *Attribute) GetDBType() string {
	return a.DBType

}
func (a *Attribute) IsPrimaryKey() bool {
	return strings.Contains(strings.ToLower(a.GetDBType()), "primary")

}
