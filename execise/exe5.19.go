package exe

import "fmt"

func weird() (ret string) {
	defer func() {
		val := recover()
		ret = "hi"
		fmt.Printf("%s\n", val)
	}()
	panic("omg")
}

func main() {
	fmt.Println(weird())
}
