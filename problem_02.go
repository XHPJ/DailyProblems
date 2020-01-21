package main

/*
Task: 	Given an array of integers, return a new array such that each element at index
		i of the new array is the product of all the numbers in the original array except the one at i.
**/

//solution without division :  no need of  special treatment for the case of an element being 0
//downside: runntime O(nˆ2)
func ProductWithoutDivision(array []int) []int {
	result := make([]int, len(array))

	/**
	This first loop  checks for a zero, in this case every other element with index != zeroIndex
	will be zero and we get a linear runtime
	*/
	for index, elem := range array {
		if elem == 0 {
			zeroIndexValue := 1

			for innerIndex, innerElem := range array {
				if innerIndex != index {
					zeroIndexValue *= innerElem
				}
			}
			result[index] = zeroIndexValue
			return result
		}
	}

	for index := range array {
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

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}
// 	fmt.Println(ProductWithoutDivision(numbers))
// }
