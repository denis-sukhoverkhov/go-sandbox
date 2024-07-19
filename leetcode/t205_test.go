package leetcode

import "testing"

func TestIsIsomorphic(t *testing.T) {

	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isIsomorphic(tt.s, tt.t); got != tt.expected {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.expected)
			}
		})
	}
}
