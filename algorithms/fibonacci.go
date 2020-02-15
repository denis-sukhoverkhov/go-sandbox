package algorithms

import "math/big"

func FibRec(n int) int {
	if n <= 1 {
		return n
	} else {
		return FibRec(n-1) + FibRec(n-2)
	}
}

func FibIter(n int) []int {
	row := []int{0, 1}
	for i := 0; i < n-2; i++ {
		value := row[len(row)-1] + row[len(row)-2]
		row = append(row, value)
	}

	return row
}

func FibIter2(n uint) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	for i := uint(1); i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}
