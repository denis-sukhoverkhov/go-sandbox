package main

import "fmt"

func utopianTree(n int32) int32 {

	current_h := int32(1)

	if n == 0 {
		return current_h
	}

	prev := 1
	for i := int32(1); i <= n; i++ {
		if prev == 1 {
			current_h *= 2
			prev = 2
		} else {
			current_h += 1
			prev = 1
		}
	}

	return current_h
}

func main() {
	fmt.Println(utopianTree(4))
}
