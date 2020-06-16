package gen

import (
	"fmt"
	"strings"
)

func templateGenerator() string {
	var code strings.Builder
	code.Grow(1000)
	fmt.Fprintf(&code, "\nname:")
	fmt.Fprintf(&code, "\nendpoints:")
	fmt.Fprintf(&code, "\n -")
	fmt.Fprintf(&code, "\n\tname:")
	fmt.Fprintf(&code, "\n\targs:")
	fmt.Fprintf(&code, "\n\toutput:")
	fmt.Fprintf(&code, "\n\ttransport:")
	fmt.Fprintf(&code, "\n\t\tmethod:")
	fmt.Fprintf(&code, "\n\t\tpath:")
	fmt.Fprintf(&code, "\nrepositroy:")
	fmt.Fprintf(&code, "\n value:")
	fmt.Fprintf(&code, "\n db:")
	fmt.Fprintf(&code, "\n\tname:")
	fmt.Fprintf(&code, "\n\taddress:")
	fmt.Fprintf(&code, "\nmodel:")
	fmt.Fprintf(&code, "\n -")
	fmt.Fprintf(&code, "\n\tname:")
	fmt.Fprintf(&code, "\n\tattr:")

	return code.String()

}
