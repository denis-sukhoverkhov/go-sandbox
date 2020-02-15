package main

import "fmt"

func reverseArray(a []int32) []int32 {
	var new_arr []int32
	for i := len(a) - 1; i >= 0; i-- {
		new_arr = append(new_arr, a[i])
	}
	return new_arr
}

func main() {
	arr := []int32{1, 4, 3, 2}
	fmt.Println(reverseArray(arr))
	fmt.Println(reverseString("cppstudio"))
}
