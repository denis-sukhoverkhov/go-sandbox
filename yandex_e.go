package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	output, _ := os.Create("output.txt")
	defer output.Close()

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	t := strings.TrimSpace(scanner.Text())
	k, _ := strconv.ParseInt(t, 0, 32)

	arr := make([]int32, 101, 101)

	for i := int64(0); i < k; i++ {
		scanner.Scan()
		line := strings.TrimSpace(scanner.Text())
		s := strings.Split(line, " ")
		sizeLine, _ := strconv.Atoi(s[0])
		if sizeLine > 0 {
			for i := 1; i < len(s); i++ {
				el, _ := strconv.Atoi(s[i])
				arr[el]++
			}
		}
	}
	for k, v := range arr {
		for i := int32(0); i < v; i++ {
			fmt.Print(k, " ")
		}
	}
}
