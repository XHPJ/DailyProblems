package main

/**
Task:	Method to find out whether k can be summed up by any two elements in the given  array/slice
*/
func FindSummands(numbers []int, k int) bool {
	for index, elem := range numbers {
		for innerindex, innerElement := range numbers {
			if innerindex == index {
				continue
			}
			if elem+innerElement == k {
				return true
			}
		}
	}
	return false
}
