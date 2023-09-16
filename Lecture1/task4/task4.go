package main

import "fmt"

//структура.тут мы храним методы сортировки
type IntSort struct {
	data []int
}

// создаем экземпляр
func NewIntSort(data []int) *IntSort {
	return &IntSort{data}
}

// бабл сорт метод
func (s *IntSort) BubbleSort() {
	n := len(s.data)
	swapping := true
	for swapping {
		swapping = false
		for i := 1; i < n; i++ {
			if s.data[i-1] > s.data[i] {
				s.data[i-1], s.data[i] = s.data[i], s.data[i-1]
				swapping = true
			}
		}
	}
}

func main() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sorter := NewIntSort(numbers)
	sorter.BubbleSort()
	fmt.Println(sorter.data)
}
