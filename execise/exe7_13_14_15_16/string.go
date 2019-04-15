package exe7_13_14_15_16

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l Literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u Unary) String() string {
	return string(u.op) + u.x.String()
}

func (b Binary) String() string {
	return fmt.Sprint(b.x.String(), string(b.op), b.y.String())
}

func (c Call) String() string {
	var buf bytes.Buffer
	buf.WriteString(c.fn)
	buf.WriteString("(")
	for i, a := range c.exprs {
		buf.WriteString(a.String())
		if i < len(c.exprs)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")
	return buf.String()
}
