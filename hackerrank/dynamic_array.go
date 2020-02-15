package main

import "fmt"

func dynamicArray(n int32, queries [][]int32) []int32 {
	lastAnswer := int32(0)

	S := make([][]int32, n)

	for i := range S {
		S[i] = []int32{}
	}

	var res []int32

	for i := int32(0); i < int32(len(queries)); i++ {
		t := queries[i][0]
		x := queries[i][1]
		y := queries[i][2]

		idx := (x ^ lastAnswer) % n
		if t == 1 {
			S[idx] = append(S[idx], y)
		} else {
			lastAnswer = S[idx][y%int32(len(S[idx]))]
			fmt.Println(lastAnswer)
			res = append(res, lastAnswer)
		}
	}

	return res
}

func main() {
	ar := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}
	fmt.Println(dynamicArray(2, ar))
}
