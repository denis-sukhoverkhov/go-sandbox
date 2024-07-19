package leetcode

import "testing"

func TestExist(t *testing.T) {

	tests := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
			true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := exist(tt.board, tt.word); got != tt.expected {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.expected)
			}
		})
	}
}
