package sort

func InsertionSort(arr []int32) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = x
	}
}
