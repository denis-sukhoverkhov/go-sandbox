package main

import "fmt"

func countingValleys(n int32, s string) int32 {
	earth := 0
	prev := 0

	ct_hills := int32(0)
	ct_valleys := int32(0)
	for _, v := range s {
		if v == 'U' {
			prev = earth
			earth++
		} else {
			prev = earth
			earth--
		}

		if earth == 0 && prev < 0 {
			ct_valleys++
		} else if earth == 0 && prev > 0 {
			ct_hills++
		}
	}

	return ct_valleys
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))
}
