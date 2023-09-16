package main

import "fmt"

func twoSum(nums []int, target int) []int {
	complementMap := map[int]int{}
	for i, n := range nums {
		c := target - n
		if j, ok := complementMap[c]; ok {
			return []int{j, i}
		}
		complementMap[n] = i
	}
	return []int{}
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	if len(result) == 2 {
		fmt.Printf("Индексы чисел, сумма которых равна %d: %d и %d\n", target, result[0], result[1])
	} else {
		fmt.Println("Не найдено.")
	}
}
