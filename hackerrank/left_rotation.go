package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []int32{1, 2, 3, 4, 5}

	n := 5

	if n > len(a) {
		n = n - len(a)
	} else if n == len(a) {
		fmt.Printf(strings.Trim(fmt.Sprint(a), "[]"))
		return
	}

	res := []int32{}

	res = append(res, a[n:]...)
	res = append(res, a[0:n]...)

	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), " "), "[]"))

}
