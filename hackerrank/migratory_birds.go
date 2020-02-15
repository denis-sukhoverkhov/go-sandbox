package main

import "fmt"

func migratoryBirds(arr []int32) int32 {

	res_types_ct := make([]int64, 5, 5)

	for _, v := range arr {
		res_types_ct[v-1]++
	}

	max := res_types_ct[0]
	cur_t := int32(0)
	for k, v := range res_types_ct {
		if v > max {
			max = v
			cur_t = int32(k)
		}
	}

	return cur_t + 1
}

func main() {
	ar := []int32{1, 1, 1, 4, 4, 4, 5, 3}
	fmt.Println(migratoryBirds(ar))
}
