package algorithms

func IsPrimeNumber(n int) bool {
	if n == 1 {
		return true
	}

	if n%2 == 0 {
		return n == 2
	}

	d := 3
	for d*d <= n && n%d != 0 {
		d += 2

	}

	return d*d > n
}
