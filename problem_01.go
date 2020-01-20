package main

import "fmt"

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

func main() {
	numbers := []int{10, 15, 3, 7}
	num := []int{96, 80, 32, 28, 15}
	fmt.Println(FindSummands(numbers, 17))
	fmt.Println(FindSummands(num, 108))
}
