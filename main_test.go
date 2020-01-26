package main

import (
	"testing"
)

//Problem 12
func TestNumberOfPossibleSteps(t *testing.T) {
	result := NumberOfPossibleSteps(4)

	if result != 5 {
		t.Errorf("Solution was incorrect, got: %d, want: %d.", result, 5)
	}
}

// func TestNumberOfStairsBySet(t *testing.T) {
// 	steps := []int{1, 2, 3}
// 	result := NumberOfStairsBySet(4, steps)
// 	expectedValues := [][]int{{1, 1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 3}, {2, 1, 1}, {2, 2}, {3, 1}}
// 	for i := 0; i < len(result); i++ {
// 		for j := 0; j < len(result[i]); j++ {
// 			if result[i][j] != expectedValues[i][j] {
// 				t.Errorf("Solution was incorrect, got: %d, want: %d.", result[i][j], expectedValues[i][j])
// 			}
// 		}
// 	}
// }
