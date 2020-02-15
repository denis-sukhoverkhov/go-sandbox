package main

import "fmt"

func minimumNumber(n int32, password string) int32 {
	number_problems := int32(0)

	if !exist_level_protect(password, "0123456789") {
		number_problems++
	}

	if !exist_level_protect(password, "abcdefghijklmnopqrstuvwxyz") {
		number_problems++
	}
	if !exist_level_protect(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		number_problems++
	}
	if !exist_level_protect(password, "!@#$%^&*()-+") {
		number_problems++
	}

	smm := int32(len(password)) + number_problems
	if smm < 6 {
		number_problems += 6 - smm
	}
	return number_problems
}

func exist_level_protect(source string, level string) bool {
	for _, v := range source {
		for _, v1 := range level {
			if v == v1 {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(minimumNumber(11, "#HackerRank"))
	fmt.Println(minimumNumber(11, "Ab1"))
	fmt.Println(minimumNumber(11, "E!%Z@"))
}
