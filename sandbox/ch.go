package sandbox

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, val := range nums {
			out <- val
		}
	}()

	return out

}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for val := range in {
			out <- val * val
		}
	}()

	return out
}
