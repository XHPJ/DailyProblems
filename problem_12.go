package main

import "fmt"

/*
TASK:		There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time. Given N, write a
			function that returns the number of unique ways you can climb the staircase. The order of the steps matters.
**/

//Just one or two steps
func NumberOfPossibleSteps(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return NumberOfPossibleSteps(n-1) + NumberOfPossibleSteps(n-2)
}

//set of possible stairs
//Note: I assume that the set is ordered and if the steps a higher than the stairs
//there is no possibility
func NumberOfStairsBySet(n int, possibleSteps []int) [][]int {
	combos := make([][]int, 0)
	if n < min(possibleSteps) {
		return combos
	}
	//buggy
	for _, steps := range possibleSteps {
		if n == steps {
			combos = append(combos, make([]int, steps))
		} else if n > steps {
			childPoss := NumberOfStairsBySet(n-steps, possibleSteps)
			for _, child := range childPoss {
				sublist := make([]int, steps)
				sublist = append(sublist, child...)
				tmp := [][]int{sublist}
				combos = append(combos, tmp...)
			}
		}
	}
	fmt.Println(combos)
	return combos
}

func min(list []int) int {
	min := MaxInt
	for _, elem := range list {
		if elem < min {
			min = elem
		}
	}
	return min
}
