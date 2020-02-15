package main

import "fmt"

func isBalanced(s string) string {

	var stack []int32

	for _, v := range s {
		switch v {
		case '{', '(', '[':
			stack = append(stack, v)
			break
		case '}':
			n := len(stack) - 1
			if len(stack) == 0 || stack[n] != '{' {
				return "NO"
			}
			stack = stack[:n] // pop
			break
		case ')':
			n := len(stack) - 1
			if len(stack) == 0 || stack[n] != '(' {
				return "NO"
			}
			stack = stack[:n] // pop
			break
		case ']':
			n := len(stack) - 1
			if len(stack) == 0 || stack[n] != '[' {
				return "NO"
			}
			stack = stack[:n] // pop
			break
		}
	}

	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println(isBalanced("{{[[(())]]}}"))
	fmt.Println(isBalanced("{[(])}"))
	fmt.Println(isBalanced("[()][{}()][](){}([{}(())([[{}]])][])[]([][])(){}{{}{[](){}}}()[]({})[{}{{}([{}][])}]"))
	fmt.Println(isBalanced("[()][{}[{}[{}]]][]{}[]{}[]{{}({}(){({{}{}[([[]][[]])()]})({}{{}})})}"))
	fmt.Println(isBalanced("{{}("))
}
