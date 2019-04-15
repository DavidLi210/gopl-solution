package exe

import "fmt"

func main() {
	fmt.Println(join(" ,", "1", "2", "3", "4"))
}

func join(sep string, strings ...string) string {
	res := ""
	for i, str := range strings {
		if i != 0 {
			res = res + sep
		}
		res = res + str
	}
	return res
}
