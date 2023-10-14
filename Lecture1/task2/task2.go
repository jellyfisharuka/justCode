package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	soz := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(strs[i]) && j < len(soz) && soz[j] == strs[i][j] {
			j++
		}
		soz = soz[:j] //slicing method
		if soz == "" {
			break
		}
	}
	return soz

}
func main() {
	input := []string{"flower", "flow", "flour"}
	result := longestCommonPrefix(input)
	fmt.Printf("Longest Common Prefix: %s\n", result)
}
