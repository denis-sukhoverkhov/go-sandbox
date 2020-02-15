package main

import "fmt"

func hackerrankInString(s string) string {
	key_word := "hackerrank"

	curStart := 0
	for i := 0; i < len(key_word); i++ {
		prevStart := curStart
		for j := curStart; j < len(s); j++ {
			if key_word[i] == s[j] {
				curStart = j + 1
				break
			}
		}

		if prevStart == curStart {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	fmt.Println(hackerrankInString("hhaacckkekraraannk"))
	fmt.Println(hackerrankInString("rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt"))

}
