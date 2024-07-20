package sandbox

import (
	"context"
	"log"
	"time"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)

	log.Printf("GEN: Was created out channel")

	go func() {
		defer close(out)

		log.Printf("GEN: Was started write to channel")
		for _, val := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- val:
				log.Printf("GEN: write")
			}
		}
	}()

	log.Printf("GEN: Was returned out channel")
	return out

}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	log.Printf("SQ: Was created out channel")

	go func() {
		defer close(out)

		log.Printf("SQ: Was started read from channel")
		for val := range in {
			select {
			case <-ctx.Done():
				return
			case out <- func() int {
				log.Printf("SQ: read")
				time.Sleep(100 * time.Millisecond)
				return val * val
			}():
			}
		}
	}()

	log.Printf("SQ: Was returned out channel")
	return out
}
