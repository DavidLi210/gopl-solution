package exe

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		strings[i] = s
		i++
	}
	return strings
}

func nonempty2(strings []string) []string {
	z := make([]string, 0, cap(strings))
	for _, s := range strings {
		if s != "" {
			z = append(z, s)
		}
	}
	return z
}

var stack = make([]int, 0)

func main() {
	ages := make(map[string]int)
	ages["David"] = 33
	ages["Bob"] = 31
	delete(ages, "Bob")
	fmt.Println(ages)
	seen := map[string]bool{}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
		}
	}
}
func remove(arr []int, i int) []int {
	copy(arr[i:], arr[i+1:])
	return arr[:len(arr)-1]
}

func remove2(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func reverse2(ptr *[8]int) {
	for i, j := 0, len(ptr)-1; i <= j; i, j = i+1, j-1 {
		temp := ptr[i]
		ptr[i] = ptr[j]
		ptr[j] = temp
	}
}

func eliminateDuplicates(ptr *[5]string) {
	for i, j := 0, 0; j < len(ptr); j++ {
		if ptr[j] != "" {
			swap(ptr, i, j)
			i++
		}
	}
}

func swap(ptr *[5]string, i int, j int) {
	var temp = ptr[i]
	ptr[i] = ptr[j]
	ptr[j] = temp
}

func rotate(nums [8]int, k int) [8]int {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			if start == current {
				break
			}
		}
	}
	return nums
}

func removeDup(b []byte) []byte {
	out := b[:0]
	for i, c := range b {
		if unicode.IsSpace(rune(c)) {
			if i > 0 && unicode.IsSpace(rune(b[i-1])) {
				continue
			} else {
				out = append(out, ' ')
			}
		} else {
			out = append(out, c)
		}
	}
	return out
}

func reverseBytes(b []byte) []byte {
	if len(b) == 1 {
		return b
	}
	_, size := utf8.DecodeRune(b)
	return append(reverseBytes(b[size:]), b[:size]...)
}
