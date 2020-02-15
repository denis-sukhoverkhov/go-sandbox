package main

import "fmt"

func superReducedString(s string) string {

	isReduced := true
	new_str := s

	for isReduced {
		isReduced = false
		temp := ""
		start_index := 0
		for i := start_index; i < len(new_str)-1; i++ {
			if new_str[i] == new_str[i+1] {
				temp = new_str[0:i]
				start_index = i
				if len(new_str)-1 >= i+1 {
					temp += new_str[i+2:]
				}
				isReduced = true
				new_str = temp
				break
			}
		}
	}

	if new_str == "" {
		return "Empty String"
	}

	return new_str
}

func main() {
	fmt.Println(superReducedString("aaabccddd"))
	fmt.Println(superReducedString("aa"))
	fmt.Println(superReducedString("baab"))

}
