package sandbox

import (
	"context"
	"time"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, val := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- val:
			}
		}
	}()

	return out

}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for val := range in {
			select {
			case <-ctx.Done():
				return
			case out <- func() int {
				time.Sleep(100 * time.Millisecond)
				return val * val
			}():
			}
		}
	}()

	return out
}
