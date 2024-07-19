package leetcode

// 1614. Maximum Nesting Depth of the Parentheses
// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

func maxDepth(s string) int {
	res := 0

	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth += 1
			res = max(depth, res)
		} else if s[i] == ')' {
			depth -= 1
		}

	}

	return res

}
