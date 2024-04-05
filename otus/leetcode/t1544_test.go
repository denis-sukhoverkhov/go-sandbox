package leetcode

import "testing"

func TestMakeGood(t *testing.T) {

	tests := []struct {
		s        string
		expected string
	}{
		{"leEeetcode", "leetcode"},
		{"abBAcC", ""},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := makeGood(tt.s); got != tt.expected {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.expected)
			}
		})
	}
}
