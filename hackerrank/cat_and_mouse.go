package main

import "math"

func catAndMouse(x int32, y int32, z int32) string {
	y_d := math.Abs(float64(z - y))
	x_d := math.Abs(float64(z - x))

	if y_d < x_d {
		return "Cat B"
	} else if x_d < y_d {
		return "Cat A"
	} else {
		return "Mouse C"
	}

}

func main() {
	println(catAndMouse(1, 3, 2))
}
