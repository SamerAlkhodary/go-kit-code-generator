package gen

import (
	"go-kit-code-generator/model"
	"strings"
)

func cacheGenerator(s model.Service) string {
	var code strings.Builder
	code.Grow(1000)

	return code.String()
}
