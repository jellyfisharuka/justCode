package main

import "fmt"

func intToRoman(num int) string {
	var res string = ""
	var romanS map[int]string = map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	var keys = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(keys); i++ {
		for num >= keys[i] {
			num -= keys[i]
			res += romanS[keys[i]]
		}
	}

	return res
}

func main() {

	num := 2000
	roman := intToRoman(num)
	fmt.Printf("The Roman numeral representation of %d is: %s\n", num, roman)
}
