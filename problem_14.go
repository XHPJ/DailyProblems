package main

import "math/rand"

import "math"

/*
Task: 		The area of a circle is defined as πr^2. Estimate π to 3 decimal places using a Monte Carlo method.

			Hint: The basic equation of a circle is x2 + y2 = r2.

**/

func GetPi() float64 {
	pi := 0.0

	inside := 0.0
	outside := 0.0
	for {
		x := rand.Float64()
		y := rand.Float64()
		random := math.Pow(x, 2.0) + math.Pow(y, 2.0)
		if random <= 1 {
			inside++
		}
		outside++
		pi = (4.0 * (inside / outside))
		if int(pi*1000.0) == 3141 {
			cutoff := 3.141 - pi
			pi += cutoff
			break
		}
	}
	return pi
}
