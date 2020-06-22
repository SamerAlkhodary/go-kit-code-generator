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
	fmt.Fprintf(&code, "\n  name:")
	fmt.Fprintf(&code, "\n  args:")
	fmt.Fprintf(&code, "\n  output:")
	fmt.Fprintf(&code, "\n  transport:")
	fmt.Fprintf(&code, "\n   method:")
	fmt.Fprintf(&code, "\n   path:")
	fmt.Fprintf(&code, "\nrepository:")
	fmt.Fprintf(&code, "\n value:")
	fmt.Fprintf(&code, "\n db:")
	fmt.Fprintf(&code, "\n  name:")
	fmt.Fprintf(&code, "\n  address:")
	fmt.Fprintf(&code, "\nmodel:")
	fmt.Fprintf(&code, "\n -")
	fmt.Fprintf(&code, "\n  name:")
	fmt.Fprintf(&code, "\n  attr:")

	return code.String()

}
