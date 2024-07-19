package leetcode

import "testing"

func TestSubarraysWithKDistinct(t *testing.T) {

	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 1, 2, 3}, 2, 7},
		{[]int{1, 2, 1, 3, 4}, 3, 3},
		{[]int{1, 2, 1, 2, 3, 3, 2, 1}, 3, 15},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := subarraysWithKDistinct(tt.nums, tt.k); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
