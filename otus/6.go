package main

import (
	"errors"
	"fmt"
	"time"
)

func Fact(n int) uint64 {
	if n <= 1 {
		return 1
	}

	return uint64(n) * Fact(n-1)
}

func Elapsed(n int) (uint64, error) {
	start := time.Now()
	value := Fact(n)

	elapsed := time.Since(start)
	if elapsed.Microseconds() > 10 {
		return 0, errors.New("big time elapsed")
	}

	return value, nil
}

func dummyJob(jobName string, n int) error {
	for i := 0; i < n; i++ {
		fmt.Printf("{%v}: %v\n", jobName, i)

		if i >= 10 {
			return errors.New("big time elapsed")
		}

	}
	return nil
}

func WrapDummyJob() error {
	return dummyJob("Test", 100)
}

func BusActions(tasks []func() error, n int, numErrors int) {
	concurrency := n
	sem := make(chan bool, concurrency)
	errChan := make(chan bool, numErrors)
	for _, f := range tasks {
		sem <- true
		go func() {
			defer func() { <-sem }()
			err := f()
			if err != nil {
				errChan <- true
			}
		}()

		if len(errChan) == cap(errChan) {
			break
		}
	}
	close(sem)

}
