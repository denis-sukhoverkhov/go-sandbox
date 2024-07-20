package sandbox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"Empty path", "", "/"},
		{"Multiple Slashes", "/////", "/"},
		{"Complex Path", "/../.././././", "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, SimplifyPath(tt.path))
		})
	}

}

func TestLoopVar(t *testing.T) {
	numbers := []*int{}
	for i := 0; i < 3; i++ {
		numbers = append(numbers, &i)
	}

	for i, num := range numbers {
		assert.Equal(t, i, *num)
	}
}

func TestSliceToPointers(t *testing.T) {
	a := []int{1, 2, 3}
	b := []*int{}

	for i := range a {
		b = append(b, &a[i])
	}

	for i, val := range a {
		assert.Equal(t, val, *b[i])
	}
}
func TestMap(t *testing.T) {
	var m map[string]int = make(map[string]int)

	fmt.Println(m["foo"])

	m["foo"] = 42

	fmt.Println(m["foo"])
}
