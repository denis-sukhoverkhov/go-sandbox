package leetcode

import "testing"

func TestCountSubarrays(t *testing.T) {

	tests := []struct {
		nums     []int
		mink     int
		maxK     int
		expected int64
	}{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
		{[]int{1, 1, 1, 1}, 1, 1, 10},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countSubarrays(tt.nums, tt.mink, tt.maxK); got != tt.expected {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.expected)
			}
		})
	}
}
