package exe7_13_14_15_16

import (
	"bufio"
	"fmt"
	"os"
)

func takeStandardInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
