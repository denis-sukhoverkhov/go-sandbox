package main

import "fmt"

// Complete the birthday function below.
func birthday(s []int32, d int32, m int32) int32 {
	ctSucc := int32(0)

	for i := 0; i < len(s); i++ {
		summ := int32(0)

		curLenBarChoc := i + int(m)
		if curLenBarChoc <= len(s) {
			for j := i; j < curLenBarChoc; j++ {
				summ += s[j]
				if summ > d {
					break
				} else {
					if summ == d && j+1 == curLenBarChoc {
						ctSucc++
						break
					}
				}
			}
		}
	}

	return ctSucc
}

func main() {
	arr := []int32{1, 2, 1, 3, 2}
	fmt.Println(birthday(arr, 3, 2))
}
