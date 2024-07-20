package sandbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	stack := &Stack{}

	stack.Push("first")
	assert.False(t, stack.IsEmpty())
	top, err := stack.Top()
	assert.Nil(t, err)
	assert.Equal(t, "first", top)
}

func TestStackPop(t *testing.T) {
	stack := &Stack{}
	assert.True(t, stack.IsEmpty())

	stack.Push("test1")
	stack.Push("test2")

	top, err := stack.Top()
	assert.Nil(t, err)
	assert.Equal(t, "test2", top)

	stack.Pop()
	top, err = stack.Top()
	assert.Nil(t, err)
	assert.Equal(t, "test1", top)
}
