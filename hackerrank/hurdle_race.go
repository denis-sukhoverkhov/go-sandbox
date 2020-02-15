package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {

	max_val := height[0]

	for _, v := range height {
		if v > max_val {
			max_val = v
		}
	}

	diff := max_val - k

	if diff <= 0 {
		return 0
	} else {
		return diff
	}
}

func main() {
	fmt.Println(hurdleRace(4, []int32{1, 6, 3, 5, 2}))
}
