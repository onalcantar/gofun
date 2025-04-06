package main

func mySqrt(x int) int {
    guess := x
	for guess*guess > x {
		guess = (guess + x / guess) / 2
	}

	return guess
}