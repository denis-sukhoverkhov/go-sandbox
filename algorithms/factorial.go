package algorithms

func FactorialRec(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * FactorialRec(n-1)
	}
}

func FactorialIter(n int) int {
	res := 1
	for i := 1; i < n+1; i++ {
		res *= i
	}

	return res
}
