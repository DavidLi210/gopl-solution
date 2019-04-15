package exe

import (
	"bufio"
	"fmt"
)

//buggy
type WordCounter int

func (w *WordCounter) Write(str string) (int, error) {
	bytearr := []byte(str)
	count, words, err := bufio.ScanWords(bytearr, false)
	if err != nil {
		fmt.Errorf("err is: %v", err)
		return 0, err
	}
	*w += WordCounter(len(words))
	return count, nil
}
func main() {
	var w WordCounter
	//w.Write("Hello World This is David")
	fmt.Println(w)
	w.Write("Test Again I am,Upset Now Haha")
	fmt.Println(w)
}
