package exe

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Triming the input")
var seq = flag.String("s", "|", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *seq))
	if !*n {
		fmt.Println()
	}
}

func delta(old, new int) int { return new - old }
