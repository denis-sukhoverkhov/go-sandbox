package main

import "fmt"

func viralAdvertising(n int32) int32 {
	comulative := 0

	shared := 5
	for i := int32(1); i <= n; i++ {
		liked := shared / 2
		comulative += liked
		shared = liked * 3
	}

	return int32(comulative)
}

func main() {
	fmt.Println(viralAdvertising(5))
}
