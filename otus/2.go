package main

import (
	"strconv"
	"strings"
)

func unpack(str string) string {
	var result string
	var prev string
	var isEscape bool = false
	for i := 0; i < len(str); i++ {
		symbol := string(str[i])

		if symbol == `\` && !isEscape {
			isEscape = true
			continue
		}

		if isEscape {
			isEscape = false
			result += symbol
			prev = symbol
			continue
		}

		r, err := strconv.Atoi(symbol)
		if err != nil {
			result += symbol
			prev = symbol
			continue
		}

		result += strings.Repeat(prev, r-1)
	}

	return result
}
