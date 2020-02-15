package main

import (
	"fmt"
)

func main() {
	var ctQueries int
	fmt.Scanf("%d", &ctQueries)

	var queries = make([][]int32, ctQueries, ctQueries)
	for i := range queries {
		queries[i] = make([]int32, 2, 2)
	}

	var a, b int32
	for i := 0; i < ctQueries; i++ {
		fmt.Scanf("%v %v", &a, &b)
		queries[i][0] = a
		queries[i][1] = b
	}

	var stack []int32

	for _, v := range queries {
		op := v[0]
		data := v[1]
		if op == 1 {
			stack = append(stack, data)
		} else if op == 2 {
			n := len(stack) - 1
			stack = stack[:n] // pop
		} else if op == 3 {
			max_val := stack[0]

			for _, v := range stack {
				if v > max_val {
					max_val = v
				}
			}
			fmt.Println(max_val)
		}
	}
}
