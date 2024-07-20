package sandbox

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSqChan(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	res := make([]int, 0, len(nums))

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*400)

	for val := range sq(ctx, gen(ctx, nums...)) {
		res = append(res, val)
	}

	assert.Equal(t, []int{1, 4, 9, 16}, res)
}

func TestGoSelect(t *testing.T) {
	var wg sync.WaitGroup
	c := make(chan string, 3)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(с chan<- string, i int, group *sync.WaitGroup) {
			defer wg.Done()
			с <- fmt.Sprintf("Goroutine %s", strconv.Itoa(i))
		}(c, i, &wg)
	}

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	close(c)
}

func TestDeadLock(t *testing.T) {
	var ch chan int = make(chan int)

	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {

		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			ch <- (idx + 1) * 2
		}(i)
	}

	go func() {
		for v := range ch {
			fmt.Println("result:", v)
			time.Sleep(2 * time.Second)
		}
	}()

	wg.Wait()
	close(ch)
}
