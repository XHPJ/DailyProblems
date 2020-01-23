package main

import "math"

/*
Task: 		Given a list of integers, write a function that returns the largest sum of non-adjacent numbers.
			Numbers can be 0 or negative. For example, [2, 4, 6, 8] should return 12, since we pick 4 and 8.
			[5, 1, 1, 5] should return 10, since we pick 5 and 5.
**/
func LargestSumOfNonAdj(values []int) int {
	firstMax := 0
	indexMax := 0
	secondMax := 0

	if len(values) < 5 {
		for index, elem := range values {
			if elem > firstMax {
				indexMax = index
				firstMax = elem
			}
		}
		for index, elem := range values {
			if math.Abs(float64(index-indexMax)) < 2 {
				continue
			}
			if elem > secondMax {
				secondMax = elem
			}
		}
		return firstMax + secondMax
	}

	for i := 0; i < len(values); i += 2 {
		firstMax += values[i]
		if i != len(values)-1 {
			secondMax += values[i+1]
		}
	}
	if firstMax > secondMax {
		return firstMax
	}
	return secondMax
}
