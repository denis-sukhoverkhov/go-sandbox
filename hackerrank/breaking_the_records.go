package main

import "fmt"

// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {

	results := make([]int32, 2, 2)

	max_val := scores[0]
	min_val := scores[0]

	for i := 1; i < len(scores); i++ {
		if scores[i] > max_val {
			results[0] += 1
			max_val = scores[i]
		} else {
			if scores[i] < min_val {
				results[1] += 1
				min_val = scores[i]
			}
		}
	}

	return results
}

func main() {
	arr := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	fmt.Println(breakingRecords(arr))
}
