package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// тест после сортировки
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	sorter := NewIntSort(input)
	sorter.BubbleSort()
	if !compareSlices(sorter.data, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorter.data)
	}
	//если слайсы вообще не менялись
	input = []int{3, 1, 4, 1, 5}
	expected = []int{1, 1, 3, 4, 5}
	sorter = NewIntSort(input)
	sorter.BubbleSort()
	if !compareSlices(sorter.data, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorter.data)
	}
	// если сортировка in reversed order
	input = []int{5, 4, 3, 2, 1}
	expected = []int{1, 2, 3, 4, 5}
	sorter = NewIntSort(input)
	sorter.BubbleSort()
	if !compareSlices(sorter.data, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorter.data)
	}

}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
