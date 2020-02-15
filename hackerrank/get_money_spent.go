package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {

	var good_sum []int32

	for _, v1 := range keyboards {
		for _, v2 := range drives {
			s := v1 + v2
			if s <= b {
				if s == b {
					return s
				}
				good_sum = append(good_sum, s)
			}
		}
	}

	if len(good_sum) == 0 {
		return -1
	}

	max_val := good_sum[0]
	for _, v := range good_sum {
		if v > max_val {
			max_val = v
		}
	}

	return max_val
}

func main() {
	kb := []int32{3, 1}
	d := []int32{5, 2, 8}

	fmt.Println(getMoneySpent(kb, d, 10))
}
