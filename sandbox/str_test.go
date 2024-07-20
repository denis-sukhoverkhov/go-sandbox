package sandbox

import (
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
