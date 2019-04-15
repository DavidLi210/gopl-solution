package exe

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "hi there"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 4)
	n, _ := b.ReadFrom(r)
	if n != 4 {
		fmt.Println("Error1")
	}
	if b.String() != "hi t" {
		fmt.Println("Error2")
	}
}
