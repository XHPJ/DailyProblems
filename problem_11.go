package main

import (
	"fmt"
	"strings"
)

/**
Task: 		Implement an autocomplete system. That is, given a query string s and a set of all possible query strings,
			return all strings in the set that have s as a prefix.
			For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
*/

//runtime O(n*m) with n as the length of the prefix and m as the length of the dictionary
func PreSelection(dic []string, prefix string) []string {
	output := []string{}

	for _, word := range dic {
		//I'm using the function from the  standard  library but there  is a self written example below
		if strings.HasPrefix(word, prefix) {
			output = append(output, word)
		}
	}
	return output
}

func main() {
	dictionary := []string{"dog", "deer", "deal"}
	predix := "de"
	fmt.Println(PreSelection(dictionary, predix))
}

//custom function
//runtime O(n) with n as the length of the prefix
func IsPrefixOf(word string, pref string) bool {
	if len(pref) > len(word) {
		return false
	}
	//empty word is always a prefix
	if pref == "" {
		return true
	}

	for i := 0; i < len(pref); i++ {
		if string(pref[i]) != string(word[i]) {
			return false
		}
	}
	return true
}
