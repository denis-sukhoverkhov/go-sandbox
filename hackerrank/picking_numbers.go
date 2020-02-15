package main

import (
	"fmt"
	"math"
)

func pickingNumbers(a []int32) int32 {
	m := make(map[int32]int)

	for _, v := range a {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	var total_max []int32
	for k1, v1 := range m {
		total_max = append(total_max, int32(v1))
		for k2, v2 := range m {
			if k1 != k2 {
				df := int32(math.Abs(float64(k1 - k2)))
				if df == 1 {
					total_max = append(total_max, int32(v1+v2))
				}
			}
		}
	}

	max := total_max[0]
	for _, v := range total_max {
		if v > max {
			max = v
		}
	}

	return int32(max)
}

func main() {

	//fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
	//fmt.Println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))
	fmt.Println(pickingNumbers([]int32{98, 3, 99, 1, 97, 2}))
}
