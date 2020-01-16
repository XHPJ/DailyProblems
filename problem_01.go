package main

import "fmt"

func FindSummands(numbers []int, k int) bool {
	for index, elem := range numbers {
		tmp := k - elem
		for innerindex, innerElement := range numbers {
			if innerindex == index {
				continue
			}
			if tmp == innerElement {
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
