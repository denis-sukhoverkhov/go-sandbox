package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	min_val_in_b := min(b)
	min_val_in_a := min(a)

	var factors_of_a []int32
	for i := min_val_in_a; i <= min_val_in_b; i++ {

		fact := 0
		for _, v := range a {
			if i%v != 0 || i < v {
				break
			} else {
				fact++
			}
		}

		if fact == len(a) {
			factors_of_a = append(factors_of_a, i)
		}
	}

	var response []int32

	for _, v := range factors_of_a {

		fact := 0
		for _, v2 := range b {
			if v2%v == 0 {
				fact++
			} else {
				break
			}
		}

		if fact == len(b) {
			response = append(response, v)
		}
	}

	return int32(len(response))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func max(a []int32) int32 {
	max_val := a[0]

	for _, v := range a {
		if v > max_val {
			max_val = v
		}
	}

	return max_val
}

func min(a []int32) int32 {
	min_val := a[0]

	for _, v := range a {
		if v < min_val {
			min_val = v
		}
	}

	return min_val
}
