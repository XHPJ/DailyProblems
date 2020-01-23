package main

import (
	"fmt"
	"time"
)

/*
Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.
**/

func SaySomething() {
	fmt.Println("I'm here!")
}

func FunctionScheduler(f func(), n int) {
	c1 := make(chan func(), 1)
	go func() {
		c1 <- f
	}()

	select {
	case <-time.After(time.Duration(n) * time.Millisecond):
		f()
	case <-time.After(time.Duration(n)*time.Millisecond + 1*time.Second):
		fmt.Println("timeout")
	}
}

func main() {
	FunctionScheduler(SaySomething, 4500)
}
