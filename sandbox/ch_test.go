package sandbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqChan(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	rnch := gen(nums...)
	sqch := sq(rnch)

	res := make([]int, len(nums))

	idx := 0
	for val := range sqch {
		res[idx] = val
		idx++
	}

	assert.Equal(t, []int{1, 4, 9, 16, 25, 36, 49}, res)
}
