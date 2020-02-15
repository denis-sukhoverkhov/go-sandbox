package main

import "fmt"

func saveThePrisoner(n int32, m int32, s int32) int32 {
	a := s + m - 1

	if a > n {
		if a%n == 0 {
			return n
		}

		return a % n
	}

	return a
}

func main() {
	fmt.Println(saveThePrisoner(5, 2, 1))
	fmt.Println(saveThePrisoner(5, 2, 2))
	fmt.Println(saveThePrisoner(7, 19, 2))
}
