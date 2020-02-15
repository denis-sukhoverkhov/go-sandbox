package main

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	first_letter := 'a'

	max_val := int32(-1)

	for _, v := range word {
		weight := v - first_letter
		if max_val < h[weight] {
			max_val = h[weight]
		}
	}

	return int32(len(word)) * 1 * max_val

}

func main() {
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc"))
}
