package exe

import (
	"fmt"
	"strings"
)

func expand(s string, f func(s string) string) string {
	replacement := f("foo")
	next := strings.Replace(s, "$foo", replacement, -1)
	return next
}

func main() {
	f := func(s string) string {
		if s == "foo" {
			return "oof"
		} else {
			return "000"
		}
	}
	fmt.Print(expand("abc$foodef e$fooc", f))
}
