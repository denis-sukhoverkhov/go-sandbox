package main

import "fmt"

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	if year <= 1917 {
		if year%4 == 0 {
			return fmt.Sprintf("12.09.%d", year)
		} else {
			return fmt.Sprintf("13.09.%d", year)
		}
	} else if year >= 1919 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			return fmt.Sprintf("12.09.%d", year)
		} else {
			return fmt.Sprintf("13.09.%d", year)
		}
	} else {
		return fmt.Sprintf("26.09.%d", year)
	}
}

func main() {
	fmt.Println(dayOfProgrammer(1800))
	fmt.Println(dayOfProgrammer(2016))
	fmt.Println(dayOfProgrammer(2017))
	fmt.Println(dayOfProgrammer(1984))
}
