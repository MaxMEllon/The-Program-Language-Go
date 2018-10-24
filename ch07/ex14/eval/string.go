package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", float64(l))
}

// A unary represents a unary operator expression, e.g., -x.
func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x.String())
}

// A binary represents a binary operator expression, e.g., x+y.
func (b binary) String() string {
	return fmt.Sprintf("%s %c %s", b.x.String(), b.op, b.y.String())
}

// A call represents a function call expression, e.g., sin(x).
func (c call) String() string {
	args := []string{}
	for _, arg := range c.args {
		args = append(args, arg.String())
	}
	argsStr := strings.Join(args, ", ")
	return fmt.Sprintf("%s(%s)", c.fn, argsStr)

}
