package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {

	m := make(map[int32]int)

	for _, v := range ar {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	summ_pairs := 0
	for _, v := range m {
		summ_pairs += v / 2
	}
	return int32(summ_pairs)
}

func main() {
	s := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(9, s))
}
