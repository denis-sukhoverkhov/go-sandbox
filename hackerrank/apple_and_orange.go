package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var calc_a []int32
	for _, v := range apples {
		calc_a = append(calc_a, v+a)
	}

	apple_ct := 0
	for _, v := range calc_a {
		if v >= s && v <= t {
			apple_ct++
		}
	}

	fmt.Println(apple_ct)

	var calc_b []int32
	for _, v := range oranges {
		calc_b = append(calc_b, v+b)
	}

	orange_ct := 0
	for _, v := range calc_b {
		if v >= s && v <= t {
			orange_ct++
		}
	}

	fmt.Println(orange_ct)

}

func main() {
	apples := []int32{-2, 2, 1}
	oranges := []int32{5, -6}

	countApplesAndOranges(7, 11, 5, 15, apples, oranges)
}
