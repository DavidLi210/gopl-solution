package exe

import (
	"bytes"
	"fmt"
	"strings"
)

type currency int

const (
	USD currency = iota
	RMB
	GBP
	EUR
)
const (
	MB = KB * 1024
	KB = 1024
	GB = 1024 * MB
)

func CountBits(x uint64) int {
	sum := 0
	i := uint(63)
	for ; i >= 0; i -= 1 {

		if ((x >> i) & 1) == 1 {
			sum++
		}
		if i == 0 {
			break
		}
	}
	return sum
}

func CountBits2(x uint64) int {
	sum := 0
	for x != 0 {
		sum++
		x = x & (x - 1)
	}
	return sum
}

func main() {
	/*//fmt.Println(CountBits2(11))
	for _, x := range "Hello World"{
		fmt.Println(string(x))
	}
	if x, y := 2, 3; x == 0 {
		fmt.Println(x)
	} else if x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}

	var num rune = 2
	num2 := rune(2)
	fmt.Println(num == num2)
	fmt.Println(-4)*/
	/*fmt.Println(11 &^ 00)
	var x uint8 = 1 << 1 | 1 << 5
	var y uint8 = 1 << 1 | 1 << 2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x & y)
	fmt.Printf("%08b\n", x | y)*/
	/*for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eA = %12.6f\n", x, math.Exp(float64(x)))
	}
	s := "Hello, 世界"
	fmt.Println(len(s)) // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i, r := range "Hello, 世界" {
		fmt.Println(r)
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}*/
	/*fmt.Println(basename("a/b/c.go") == basename2("a/b/c.go"))
	fmt.Println(intsToString([]int{1, 2, 3, 4}))*/
	//fmt.Println(comma("asdqwezxc"))
	//fmt.Println(anagram("clouds", "dlcou"))
	symbols := [...]string{USD: "$", EUR: "&", RMB: "y", GBP: "g"}
	fmt.Println(RMB, symbols[2])
	nums1 := [2]int{1, 2}
	nums2 := [...]int{1, 2}
	fmt.Println(nums1 == nums2)
	fmt.Printf("Type is %T. Then is %T.", nums1, nums1[1:])
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		//buf.WriteString(string(v))
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		return s
	}

	for i := 0; i < len(s); i++ {
		buf.WriteRune(rune(s[i]))
		if i == 3 {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	arr1 := [256]int{}
	arr2 := [256]int{}
	for ch := range s1 {
		arr1[ch]++
	}

	for ch := range s2 {
		arr2[ch]++
	}

	for i := 0; i < 256; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
