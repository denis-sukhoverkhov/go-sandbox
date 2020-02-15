package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {

	ct_upper := int32(1)
	for _, v := range s {
		if unicode.IsUpper(v) {
			ct_upper++
		}
	}

	return ct_upper
}

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}
