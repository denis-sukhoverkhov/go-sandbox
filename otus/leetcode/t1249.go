package leetcode

// 1249. Minimum Remove to Make Valid Parentheses
// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

func minRemoveToMakeValid(s string) string {
	tmp := []rune{}
	ct := 0

	for _, c := range s {
		if ct == 0 && c == ')' {
			continue
		}

		if c == ')' {
			ct -= 1
		}

		if c == '(' {
			ct += 1
		}
		tmp = append(tmp, c)
	}

	reversed := []rune{}
	for i := len(tmp) - 1; i >= 0; i-- {
		if ct > 0 && tmp[i] == '(' {
			ct -= 1
		} else {
			reversed = append(reversed, tmp[i])
		}
	}

	res := []rune{}

	for i := len(reversed) - 1; i >= 0; i-- {
		res = append(res, reversed[i])
	}

	return string(res)
}
