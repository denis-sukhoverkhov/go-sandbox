package leetcode

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {

	tests := []struct {
		s        string
		expected string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.s); got != tt.expected {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}
