package main

import (
	"fmt"
	"unicode"
)

func pangrams(s string) string {
	m := map[rune]bool{}

	for _, v := range s {
		if v != ' ' {
			lower_case := unicode.ToLower(v)
			m[lower_case] = true
		}
	}

	if len(m) == 26 {
		return "pangram"
	} else {
		return "not pangram"
	}

}

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))
}
