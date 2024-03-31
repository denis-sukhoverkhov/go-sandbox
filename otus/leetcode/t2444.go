package leetcode

func countSubarrays(nums []int, minK int, maxK int) int64 {

	var res int64 = 0

	start := 0
	last_mink_idx := -1
	last_maxk_idx := -1
	for r := 0; r < len(nums); r++ {
		if nums[r] > maxK || nums[r] < minK {
			start = r + 1
			continue
		}

		if nums[r] == minK {
			last_mink_idx = r
		}

		if nums[r] == maxK {
			last_maxk_idx = r
		}

		if last_maxk_idx >= start && last_mink_idx >= start {
			min_idx := min(last_maxk_idx, last_mink_idx)
			res += int64(min_idx) - int64(start) + 1
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
