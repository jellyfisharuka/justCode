package main
import "fmt"
func comparingSlice(sl1 []int, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	} else {
		for i := 0; i < len(sl1); i++ {
			if sl1[i] != sl2[i] {
				return false
			}
		}
		return true
	}
}
func main() {
	slice1 := []int{2, 2, 2}
	slice2 := []int{2, 2, 2}
	slice3 := []int{2, 3}
	fmt.Print("comparing result(slice1, slice2): ", comparingSlice(slice1, slice2), " \ncomparing result(slice2, slice3): ", comparingSlice(slice2, slice3), " \ncomparing result(slice1, slice3): ", comparingSlice(slice1, slice3))
}

