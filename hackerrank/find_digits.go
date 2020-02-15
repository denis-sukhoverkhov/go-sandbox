package main

import "fmt"

func findDigits(n int32) int32 {
	mp := n

	dig_arr := []int32{}

	for mp != 0 {
		new_dig := mp % 10
		mp /= 10

		if new_dig != 0 {
			dig_arr = append(dig_arr, new_dig)
		}
	}

	res_ct := int32(0)
	for _, v := range dig_arr {
		if n%v == 0 {
			res_ct++
		}
	}
	return res_ct
}

func main() {
	fmt.Println(findDigits(1012))
}
