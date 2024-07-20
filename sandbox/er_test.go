package sandbox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyError(t *testing.T) {

	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{"Nil error", (*myError)(nil), "<nil>"},
		{"My error", myError{code: 500}, "My error 500"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, fmt.Sprintf("%v", tt.err))
		})
	}

}

func TestMyError2(t *testing.T) {

	var e myError = myError{code: 500}

	assert.Equal(t, "My error 500", fmt.Sprintf("%v", e))
}
