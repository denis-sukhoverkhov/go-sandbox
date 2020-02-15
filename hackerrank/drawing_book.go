package main

import "fmt"

func pageCount(n int32, p int32) int32 {

	if p == 1 || p == n {
		return 0
	}

	count_from_start := p / 2
	count_from_last := (n - p) / 2

	if count_from_last == 0 && n-p == 1 {
		return 1
	}

	if count_from_start > count_from_last {
		return count_from_last
	} else {
		return count_from_start
	}

}

func main() {

	fmt.Println(pageCount(6, 5))
}
