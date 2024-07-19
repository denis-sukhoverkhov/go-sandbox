package leetcode

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.expected {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.expected)
			}
		})
	}
}
