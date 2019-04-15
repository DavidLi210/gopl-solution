package exe

import (
	"bytes"
	"fmt"
)

var root *tree

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	buf := &bytes.Buffer{}
	values := make([]int, 0)
	values = appendValues(values, t)
	fmt.Fprintf(buf, "[%d", values[0])
	for _, num := range values[1:] {
		fmt.Fprintf(buf, " %d", num)
	}
	fmt.Fprintf(buf, "]")
	return buf.String()
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

type Point struct {
	X, Y int
}

func Sort(values []int) {
	for _, v := range values {
		root = add(root, v)
	}
	values = appendValues(values, root)
	fmt.Println(values)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}

	if root.value < value {
		root.right = add(root.right, value)
	} else {
		root.left = add(root.left, value)
	}
	return root
}

func main() {
	arr := []int{1, 5, 7, 2, 3, 10, 8, 0}
	Sort(arr)
	fmt.Println(root.String())
}
