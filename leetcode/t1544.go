package leetcode

import "unicode"

// 1544. Make The String Great
// https://leetcode.com/problems/make-the-string-great/

func makeGood(s string) string {

	stack := []rune{'0'}
	for _, current := range s {
		head := stack[len(stack)-1]

		head_lower := unicode.ToLower(stack[len(stack)-1])
		current_lower := unicode.ToLower(current)

		if head_lower == current_lower && (unicode.IsLower(head) && unicode.IsUpper(current) || unicode.IsLower(current) && unicode.IsUpper(head)) {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, current)
	}

	return string(stack[1:])
}
