package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Task:			Given a stream of elements too large to store in memory, pick a random element from
				the stream with uniform probability.
*/

func RandomUniformSelection() {
	stream := make(chan int)
	quit := make(chan int)

	//Simulate infinite stream
	go func() {
		for {
			random := rand.Intn(100)
			stream <- random
		}
	}()

	var sample [10000]int

	//closes the channel after 5 Seconds
	go func() {
		time.AfterFunc(5000*time.Millisecond, func() { quit <- 0 })
	}()

	for {
		select {
		case num := <-stream:
			randomIndex := rand.Intn(len(sample))
			sample[randomIndex] = num
		case <-quit:
			countmap := make(map[int]int)
			for _, elem := range sample {
				countmap[elem]++
			}
			for k, v := range countmap {
				fmt.Printf("key[%d] value[%d]\n", k, v)
			}
			return
		}
	}

}
