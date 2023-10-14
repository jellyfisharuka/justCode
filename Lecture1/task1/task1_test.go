package main
import (
	"testing"
)

func compareSlices(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expect := []int{0, 1}
	result := twoSum(nums, target)
	if !compareSlices(result, expect) {
		t.Errorf("Ожидалось %v, получено %v", expect, result)
	}

	nums = []int{1, 2, 3, 4, 5}
	target = 10
	expect = []int{}
	result = twoSum(nums, target)
	if !compareSlices(result, expect) {
		t.Errorf("Ожидалось %v, получено %v", expect, result)
	}
}
