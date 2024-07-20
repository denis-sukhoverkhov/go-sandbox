package sandbox

import "testing"

func TestWorkerPool(t *testing.T) {
	pool := NewWorkerPool(10)

	for i := 0; i < 100; i++ {
		pool.TaskQueue <- Task{ID: i + 1}
	}

	close(pool.TaskQueue)

	pool.Wait()
}
