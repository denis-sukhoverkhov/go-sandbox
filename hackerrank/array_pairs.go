package main

//
//import "fmt"
//
//func solve(arr []int32) int64 {
//	max_val, index := max_value_and_index(arr)
//	if len(arr) <= 1 {
//		return 0
//	}
//	return solve(arr[:index]) + solve(arr[index:])
//}
//
//func max_value_and_index(a []int32) (int32, int32) {
//	index := 0
//	max_val := a[index]
//	for k, v := range a {
//		if v > max_val {
//			max_val = v
//			index = k
//		}
//	}
//
//	return max_val, int32(index)
//}
//
//func main()  {
//	fmt.Print(solve([]int32{1, 1, 2, 4, 2}))
//}

func main() {
	v := 5
	p := &v
	println(*p)

	changePointer(p)
	println(*p)
}

func changePointer(p *int) {
	//v := 3
	p = 3
}
