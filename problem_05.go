package main

import "fmt"

/*
cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair.
For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

def cons(a, b):
    def pair(f):
        return f(a, b)
	return pair


this is actually not the way you would do that in golang
**/

// here an example of getters and setters in golang

type Pair struct {
	first  int
	second int
}

//setter with the keyword "Set...()" as it is common in other programming languages
func (pair *Pair) SetFirst(newFirst int) {
	pair.first = newFirst
}

//getter without the "get...()" just the name of the attribute
func (pair *Pair) First() int {
	return pair.first
}

func (pair *Pair) SetSecond(newSecond int) {
	pair.second = newSecond
}

func (pair *Pair) Second() int {
	return pair.second
}

func main() {
	duo := &Pair{first: 4, second: 2}

	fmt.Println(duo.First())
	fmt.Println(duo.Second())
	fmt.Println(duo.First(), duo.Second())
}
