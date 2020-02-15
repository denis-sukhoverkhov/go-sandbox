package main

import (
	"fmt"
	"math"
	"strconv"
)

func beautifulDays(i int32, j int32, k int32) int32 {

	ct := int32(0)
	for ix := i; ix <= j; ix++ {
		s := strconv.Itoa(int(ix))
		s = reverseString(s)
		reversed_num, _ := strconv.Atoi(s)
		val := int32(math.Abs(float64(ix - int32(reversed_num))))
		if val%k == 0 {
			ct++
		}
	}

	return ct
}

func reverseString(s string) string {
	if len(s) == 0 {
		return ""
	}
	n := len(s) - 1
	return string(s[n]) + reverseString(s[:n])
}

func main() {
	//fmt.Println(7%3)
	fmt.Println(beautifulDays(20, 23, 6))
	fmt.Println(beautifulDays(13, 45, 3))
}
