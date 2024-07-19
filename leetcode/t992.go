package leetcode

// 992. Subarrays with K Different Integers
// https://leetcode.com/problems/subarrays-with-k-different-integers/

func subarraysWithKDistinct(nums []int, k int) int {

	res := 0
	l_far := 0
	l_near := 0
	m := make(map[int]int)
	for r := 0; r < len(nums); r++ {
		m[nums[r]] += 1

		for len(m) > k {
			m[nums[l_near]] -= 1

			if m[nums[l_near]] == 0 {
				delete(m, nums[l_near])
			}

			l_near += 1
			l_far = l_near
		}

		for m[nums[l_near]] > 1 {
			m[nums[l_near]] -= 1
			l_near += 1
		}

		if len(m) == k {
			res += l_near - l_far + 1
		}

	}

	return res

}
