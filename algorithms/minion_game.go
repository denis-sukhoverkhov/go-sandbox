package algorithms

import (
	"fmt"
	"unicode/utf8"
)

func MinionGame(str string) {
	vowels := []uint8{'A', 'E', 'I', 'O', 'U'}
	kevsc, stusc := 0, 0

	ln := utf8.RuneCount([]byte(str))

	for i := 0; i < ln-1; i++ {
		if ExistsInArrayString(str[i], vowels) {
			kevsc += (ln - i)
		} else {
			stusc += (ln - i)
		}
	}

	if kevsc > stusc {
		fmt.Printf("Kevin %v\n", kevsc)
	} else {
		fmt.Printf("Stuart %v\n", stusc)
	}
}

func ExistsInArrayString(value uint8, slice []uint8) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}
