package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	summ := int32(0)

	for ki, v := range bill {
		if ki != int(k) {
			summ += v
		}
	}

	res := summ / 2

	if res == b {
		fmt.Println("Bon Appetit")
	} else if res < b {
		fmt.Println(b - res)
	}

}

func main() {
	bill := []int32{3, 10, 2, 9}
	bonAppetit(bill, 1, 12)
}
