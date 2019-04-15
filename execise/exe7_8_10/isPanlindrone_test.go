package exe7_8_10

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	if IsPalindrome(sort.IntSlice(arr1)) {
		t.Fail()
	}
	arr2 := []int{1, 2, 2, 1}

	if !IsPalindrome(sort.IntSlice(arr2)) {
		t.Fail()
	}
}
