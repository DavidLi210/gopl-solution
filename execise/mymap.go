package exe

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadString('\n')
		fmt.Println(line)
		if line == "end\n" {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount", err)
		}
		arr := strings.Split(line, " ")
		for _, word := range arr {
			counts[word]++
		}
	}
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
}
