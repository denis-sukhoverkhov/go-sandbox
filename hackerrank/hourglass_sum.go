package main

import "fmt"

func hourglassSum(arr [][]int32) int32 {
	var summList []int32

	rows := len(arr)
	colls := len(arr[0])

	for i := 0; i <= rows-3; i++ {
		for j := 0; j <= colls-3; j++ {
			s := sum(arr[i][j:j+3]) + sum(arr[i+1][j+1:j+2]) + sum(arr[i+2][j:j+3])
			summList = append(summList, s)
		}
	}

	max_val := summList[0]
	for _, v := range summList {
		if v > max_val {
			max_val = v
		}
	}

	return max_val
}

func sum(ar []int32) int32 {
	s := int32(0)
	for _, v := range ar {
		s += int32(v)
	}

	return s
}

func main() {
	arr := [][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	fmt.Println(hourglassSum(arr))
}
