package sandbox

import (
	"context"
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
