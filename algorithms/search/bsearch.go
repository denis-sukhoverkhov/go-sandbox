package search

func BSearch(numbers []int, val int) int {
	first := 0
	last := len(numbers) - 1
	var index int = -1

L:
	for first <= last {
		mid := (first + last) / 2

		switch {
		case numbers[mid] == val:
			index = mid
			break L
		case numbers[mid] > val:
			last = mid - 1
		default:
			first = mid + 1
		}
	}

	return index
}
