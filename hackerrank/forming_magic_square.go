package main

import (
	"fmt"
	"math"
)

func formingMagicSquare(s [][]int32) int32 {
	//the_same_elements := make(map[int32]int32, 9)

	arr := createMagicSquare(3)

	var transp_arr = make([][]int32, 3)
	for i := range transp_arr {
		transp_arr[i] = make([]int32, 3)
	}

	for i, _ := range arr {
		for j, _ := range arr {
			transp_arr[j][i] = arr[i][j]
		}
	}

	var arr_new [][]int32
	for i, _ := range arr {
		var tmp []int32
		for j := len(arr[i]) - 1; j >= 0; j-- {
			tmp = append(tmp, int32(arr[i][j]))
		}
		arr_new = append(arr_new, tmp)
	}

	var arr_new_tr [][]int32
	for i, _ := range transp_arr {
		var tmp []int32
		for j := len(transp_arr[i]) - 1; j >= 0; j-- {
			tmp = append(tmp, int32(transp_arr[i][j]))
		}
		arr_new_tr = append(arr_new_tr, tmp)
	}

	var arr_of_arrs [][][]int32

	//arr_of_arrs := [][][]int32{
	//	{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
	//	{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	//	{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
	//	{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
	//	{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
	//	{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
	//	{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
	//	{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	//}

	arr_of_arrs = append(arr_of_arrs, arr)
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr[1], arr[2], arr[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr[2], arr[1], arr[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr[1], arr[0], arr[2]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr[2], arr[0], arr[1]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr[0], arr[2], arr[1]})
	//
	//arr_of_arrs = append(arr_of_arrs, arr_new)
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new[1], arr_new[2], arr_new[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new[2], arr_new[1], arr_new[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new[1], arr_new[0], arr_new[2]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new[2], arr_new[0], arr_new[1]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new[0], arr_new[2], arr_new[1]})

	//arr_of_arrs = append(arr_of_arrs, transp_arr)
	//arr_of_arrs = append(arr_of_arrs, [][]int32{transp_arr[1], transp_arr[2], transp_arr[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{transp_arr[2], transp_arr[1], transp_arr[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{transp_arr[1], transp_arr[0], transp_arr[2]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{transp_arr[2], transp_arr[0], transp_arr[1]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{transp_arr[0], transp_arr[2], transp_arr[1]})
	//
	//arr_of_arrs = append(arr_of_arrs, arr_new_tr)
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new_tr[1], arr_new_tr[2], arr_new_tr[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new_tr[2], arr_new_tr[1], arr_new_tr[0]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new_tr[1], arr_new_tr[0], arr_new_tr[2]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new_tr[2], arr_new_tr[0], arr_new_tr[1]})
	//arr_of_arrs = append(arr_of_arrs, [][]int32{arr_new_tr[0], arr_new_tr[2], arr_new_tr[1]})

	for _, v := range arr_of_arrs {
		fmt.Println(v)
	}

	var totals []int32
	for i := 0; i < len(arr_of_arrs); i++ {
		a := arr_of_arrs[i]
		total := int32(0)
		for k1, v1 := range a {
			for k2, v2 := range v1 {
				if s[k1][k2] != v2 {
					total += int32(math.Abs(float64(s[k1][k2] - v2)))
				}
			}
		}
		totals = append(totals, total)
	}

	min_val := totals[0]

	for _, v := range totals {
		if v < min_val {
			min_val = v
		}
	}

	return min_val
}

func createMagicSquare(n int) [][]int32 {
	arr := make([][]int32, n)

	for i := range arr {
		arr[i] = make([]int32, n, n)
	}

	i, j := 0, 1
	newI, newJ := i-1, j-1

	for v := 1; v <= n*n; v++ {
		newI, newJ = (newI+1)%n, (newJ+1)%n

		tmpJ := newJ
		if newJ < 0 {
			tmpJ = newJ + n
		}

		tmpI := newI
		if newI > n-1 {
			tmpI = newI - n
		}

		arr[tmpI][tmpJ] = int32(v)

		if v%n == 0 {
			i++
			j--
			newI, newJ = i-1, j-1
		}
	}

	return arr
}

func main() {
	s := [][]int32{
		//{4, 8, 2},
		//{4, 5, 7},
		//{6, 1, 6},

		//{4, 9, 2},
		//{3, 5, 7},
		//{8, 1, 5},

		{4, 5, 8},
		{2, 4, 1},
		{1, 9, 7},
	}
	fmt.Println(formingMagicSquare(s))
}
