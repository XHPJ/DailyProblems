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

//Problem 13
func TestLongestSubstring(t *testing.T) {
	result := LongestSubstring(2, "abcbbbabbcbbadd")

	if result != "bbbabb" {
		t.Errorf("Solution was incorrect, got: %v, want: %v.", result, "bcb")
	}
}

//Problem 14
func TestGetPi(t *testing.T) {
	result := GetPi()

	if result != 3.141 {
		t.Errorf("Solution was incorrect, got: %f, want: %f.", result, 3.141)
	}

}
