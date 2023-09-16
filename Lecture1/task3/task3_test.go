package main

import (
	"testing"
)

func TestComparingSlice(t *testing.T) {
	t.Run("Equal Slices", func(t *testing.T) {
		slice1 := []int{2, 2, 2}
		slice2 := []int{2, 2, 2}
		res := compareSlice(slice1, slice2)
		if !res {
			t.Errorf("Expected comparingSlice(slice1, slice2) to be true, but got false")
		}
	})

	t.Run("Different Length Slices", func(t *testing.T) {
		slice1 := []int{2, 2, 2}
		slice3 := []int{2, 3}
		res := compareSlice(slice1, slice3)
		if res {
			t.Errorf("Expected comparingSlice(slice1, slice3) to be false, but got true")
		}
	})

	t.Run("Different Slices", func(t *testing.T) {
		slice2 := []int{2, 2, 2}
		slice3 := []int{2, 3}
		res := compareSlice(slice2, slice3)
		if res {
			t.Errorf("Expected comparingSlice(slice2, slice3) to be false, but got true")
		}
	})
}
