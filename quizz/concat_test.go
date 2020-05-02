package quizz

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func concat1(n int) time.Duration {
	start := time.Now()
	s := ""
	for i := 0; i < n; i++ {
		s += "a"
	}

	dur := time.Since(start)
	fmt.Println("Concat1: %V", dur)
	return dur
}

func concat2(n int) time.Duration {
	start := time.Now()
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString("a")
	}

	dur := time.Since(start)
	fmt.Println("Concat2: %V", dur)
	return dur
}

func concat3(n int) time.Duration {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = "a"
	}
	start := time.Now()
	strings.Join(s, "")
	dur := time.Since(start)
	fmt.Println("Concat3: %V", dur)
	return dur
}

func TestConcat1(t *testing.T) {
	t.Run("concat1", func(t *testing.T) {

		//assert.True(t, concat1(1000) > concat2(1000))
		concat1(1000)
	})
}

func TestConcat2(t *testing.T) {

	t.Run("concat2", func(t *testing.T) {

		//assert.True(t, concat2(10000) < concat3(10000))
		concat2(10000)
	})

}

func TestConcat3(t *testing.T) {

	t.Run("slowest concat method", func(t *testing.T) {

		//assert.True(t, concat1(1000) > concat3(1000))
		concat3(10000)
	})
}
