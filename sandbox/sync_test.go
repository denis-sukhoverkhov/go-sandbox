package sandbox

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	m := make(map[int]int)

	mux := sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			mux.Lock()
			m[key] = key * key
			mux.Unlock()
		}(i)
	}

	wg.Wait()

	for key, val := range m {
		fmt.Printf("Key: %d, Value: %d\n", key, val)
	}

}
