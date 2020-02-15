package algorithms

import (
	"bytes"
	"unicode/utf8"
)

func ReverseString(s string) string {
	var buffer bytes.Buffer

	for i := utf8.RuneCount([]byte(s)) - 1; i >= 0; i-- {
		buffer.WriteByte(s[i])
	}

	return buffer.String()
}

func ReverseString2(s string) string {
	i := 0
	j := utf8.RuneCount([]byte(s)) - 1

	var newString = s

	for i <= j {
		i_ := string(newString[i])
		j_ := string(newString[j])

		sl1 := newString[:i]
		sl2 := newString[i+1 : j]
		sl4 := newString[j+1:]
		newString = sl1 + j_ + sl2 + i_ + sl4
		i += 1
		j -= 1
	}

	return newString
}
