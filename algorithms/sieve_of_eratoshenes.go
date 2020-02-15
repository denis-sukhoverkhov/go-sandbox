package algorithms

import "fmt"

func sieve_of_eratoshenes(n int32) {
	numbers := make([]bool, n, n)

	for k, _ := range numbers {
		numbers[k] = true
	}
	numbers[0] = false
	numbers[1] = false

	p := int32(2)
	for p*p <= n {
		if numbers[p] {
			for i := p * p; i < n; i += p {
				numbers[i] = false
			}
		}
		p++
	}

	for k, v := range numbers {
		if v == true {
			fmt.Println(k)
		}
	}
}
