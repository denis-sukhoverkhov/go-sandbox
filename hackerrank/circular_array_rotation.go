package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {

	new2_arr := []int32{}
	rot := k % int32(len(a))
	for _, v := range queries {
		i := v - rot
		if v-rot < 0 {
			i += int32(len(a))
		}
		new2_arr = append(new2_arr, a[i])
	}

	return new2_arr
}

func main() {
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4, 5}, 3, []int32{0, 2, 4}))
}
