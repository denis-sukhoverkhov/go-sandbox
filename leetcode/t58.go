package leetcode

// 58. Length of Last Word
// https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	res := 0

	f := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if f {
				break
			}
		} else {
			res += 1
			f = true
		}

	}

	return res

}
