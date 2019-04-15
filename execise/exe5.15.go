package exe

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{11, 1, 3, 5, 10, 20, 99, 111}
	fmt.Printf("min: %d", max(arr...))
	fmt.Printf("min: %d", min(arr...))
}

func max(nums ...int) int {
	res := math.MinInt32
	for _, num := range nums {
		res = Max(res, num)
	}
	return res
}

func min(nums ...int) int {
	res := math.MaxInt32
	for _, num := range nums {
		res = Min(res, num)
	}
	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
