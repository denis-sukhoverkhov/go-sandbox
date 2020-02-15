package main

import "fmt"

func angryProfessor(k int32, a []int32) string {

	ct := int32(0)

	for _, v := range a {
		if v <= 0 {
			ct++
		}
	}

	if ct < k {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(angryProfessor(2, []int32{0, -1, 2, 1}))
}
