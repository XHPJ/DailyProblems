package main

import (
	"strconv"
)

/*
TASK:		Given the mapping a = 1, b = 2, ... z = 26, and an encoded message,
			count the number of ways it can be decoded.
			For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
			You can assume that the messages are decodable. For example, '001' is not allowed.
**/

func NumOfPossibleDecodes(input string) int {
	//convert  input string into a number and check for error
	num, err := strconv.Atoi(input)
	middle := len(input) / 2

	if err != nil {
		panic("Invalid input")
	}
	// check recursivly if the substring (always in half) can be split up
	// into  smaller numbers until we have a num < 10
	if num == 0 {
		return 0
	} else if num > 0 && num < 10 {
		return 1
	} else if num >= 10 && num < 27 {
		return 1 + NumOfPossibleDecodes(input[0:middle]) + NumOfPossibleDecodes(input[middle:len(input)])
	} else {
		return NumOfPossibleDecodes(input[0:middle]) + NumOfPossibleDecodes(input[middle:len(input)])
	}
}
