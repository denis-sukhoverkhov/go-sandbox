package sandbox

import "strings"

func SimplifyPath(path string) string {
	components := strings.Split(path, "/")

	var stack Stack

	for _, component := range components {
		if component == "" || component == "." {
			continue
		}

		if component == ".." {
			stack.Pop()
		} else {
			stack.Push(component)
		}
	}

	return "/" + strings.Join(stack.Items(), "/")
}
