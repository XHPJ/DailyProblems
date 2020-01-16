package main

import "fmt"

//solution without division to be  preferred  due to  the need of  special treatment
//for the case of an element being 0
func ProductWithoutDivision(array []int) []int {
	var result = make([]int, len(array))

	for index, _ := range array {
		tmp := 1
		for innerIndex, elem := range array {
			if innerIndex != index {
				tmp *= elem
			}
		}
		result[index] = tmp
		tmp = 1
	}

	return result
}

/**


 */
func main() {
	numbers := []int{0, 2, 3, 4, 5, 6, 0, 8, 9}
	// num := []int{1, 2, 3}
	fmt.Println(ProductWithoutDivision(numbers))
}
