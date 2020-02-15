package algorithms

func FastPowerRec(base, power int) int {
	if power == 0 {
		return 1
	}

	p := FastPowerRec(base, int(power/2))
	p *= p

	if power%2 > 0 {
		p *= base
	}

	return p
}

func FastPowerIter(base, power int) int {
	result := 1

	for power > 0 {
		if power%2 > 0 {
			power -= 1
			result *= base
		}

		if power != 0 {
			power = power / 2
			base *= base
		}
	}

	return result
}
