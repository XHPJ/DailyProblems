package main

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

/*
Task: 	Given an array of integers, find the first missing positive integer in linear time
		and constant space. In other words, find the lowest positive integer that does not exist in the array.
		The array can contain duplicates and negative numbers as well.
**/

func FirstMissingPositive(array []int) int {
	min := MaxInt
	missing := 0

	for _, elem := range array {
		if elem < min && elem >= 1 {
			//in case that the delta is larger than 1 then there is a larger gap then just one
			if (min-elem) > 1 && missing != elem {
				missing = elem + 1
			}
			min = elem
		} else if min+1 == elem {
			missing = elem + 1
		}
	}
	if min > 1 && min > missing {
		return min - 1
	}
	return missing
}

// func main() {
// 	numbers := []int{3, 4, -1, 1}

// 	fmt.Println(FirstMissingPositive(numbers))
// }
