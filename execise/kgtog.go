package exe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func init() {
	fmt.Println(1)
}

func init() {
	fmt.Println(2)
}

func init() {
	fmt.Println(4)
}

func init() {
	fmt.Println(3)
}
func main() {
	if len(os.Args) <= 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			strs := scanner.Text()
			fmt.Println(strs)
			if strs == "exit" {
				return
			}
		}

		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
		}
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "kgtg: %v\n\r", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "corresponding g = %v", t*1000)
	}
}
