package algorithms

import (
	"fmt"
	"math"
)

func SquareRoot(number float64) string {
	num_iterations := 0
	epicilon := 0.0001
	ans := 0.0

	for math.Abs(math.Pow(ans, 2)-number) > epicilon {
		ans += 0.00001
		num_iterations += 1
	}

	return fmt.Sprintf("root = %v, iteration = %v", ans, num_iterations)
}

func BSquareRoot(number float64) string {
	num_iterations := 0
	epicilon := 0.0001
	low := 0.0
	high := number
	guess := (low + high) / 2

	for math.Abs(math.Pow(guess, 2)-number) > epicilon {
		if math.Pow(guess, 2) < number {
			low = guess
		} else {
			high = guess
		}

		guess = (low + high) / 2
		num_iterations += 1
	}

	return fmt.Sprintf("root = %v, iteration = %v", guess, num_iterations)
}
