package exe

import "fmt"

func main() {
	var y int
	var pointerToY *int
	var pointerToPointerToInt **int

	y = 10
	pointerToY = &y
	pointerToPointerToInt = &pointerToY

	fmt.Println("y: ", y)
	fmt.Println("pointerToY: ", pointerToY)
	fmt.Println("pointerToPointerToInt: ", pointerToPointerToInt)

	fmt.Println("&y: ", &y)                                         // address of y
	fmt.Println("&pointerToY: ", &pointerToY)                       // address of pointerToY
	fmt.Println("&pointerToPointerToInt: ", &pointerToPointerToInt) // address of pointerToPointerToInt

	// fmt.Println(*y) throws an error because
	// you can't redirect without an address..
	// y only has int value of 10
	fmt.Println("*pointerToY: ", *pointerToY)                       // gives the value of y
	fmt.Println("*pointerToPointerToInt: ", *pointerToPointerToInt) // gives the value of pointerToY which is the address of y

	fmt.Println("**pointerToPointerToInt: ", **pointerToPointerToInt) // this gives 10, because we are redirecting twice to get y

	if pointerToY == *pointerToPointerToInt {
		fmt.Println("'pointerToY == *pointerToPointerToInt' are the same!")
	}

	if pointerToY == &y {
		fmt.Println("'pointerToY == &y' are the same!")
	}

	if &pointerToY == pointerToPointerToInt {
		fmt.Println("'&pointerToY == pointerToPointerToInt' are the same!")
	}

	if y == **pointerToPointerToInt {
		fmt.Println("'y == **pointerToPointerToInt' are the same!")
	}

	if pointerToY == *pointerToPointerToInt {
		fmt.Println("'pointerToY == *pointerToPointerToInt' are the same!")
	}

}
