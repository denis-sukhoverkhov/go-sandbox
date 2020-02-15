package main

import "fmt"

func arrayManipulation(n int32, queries [][]int32) int64 {
	zerroArray := make([]int64, n, n)

	for _, q := range queries {
		zerroArray[q[0]-1] += int64(q[2])

		if int(q[1]) < len(zerroArray) {
			zerroArray[q[1]] -= int64(q[2])
		}
	}

	max_val, x := int64(0), int64(0)

	for _, v := range zerroArray {
		x += v
		if max_val < x {
			max_val = x
		}
	}
	return max_val
}

func main() {
	arr := [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}
	fmt.Print(arrayManipulation(5, arr))
}
