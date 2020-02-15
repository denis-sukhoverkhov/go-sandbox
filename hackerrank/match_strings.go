package main

import "fmt"

func matchingStrings(strings []string, queries []string) []int32 {
	var total_matches []int32

	for i := 0; i < len(queries); i++ {
		matches := 0
		for _, v := range strings {
			if queries[i] == v {
				matches++
			}
		}
		total_matches = append(total_matches, int32(matches))
	}

	return total_matches
}

func main() {
	ar := []string{
		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
		"na",
		"asdjf",
		"na",
		"basdn",
		"sdaklfj",
		"asdjf",
	}

	q := []string{
		"abcde",
		"sdaklfj",
		"asdjf",
		"na",
		"basdn",
	}
	fmt.Println(matchingStrings(ar, q))
}
