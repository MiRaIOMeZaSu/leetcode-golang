package main

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findLHS(nums []int) int {
	// 使用窗口法?
	left := 0
	right := 1
	ans := 0
	min := min(nums[left], nums[right])
	max := max(nums[left], nums[right])
	for right < len(nums) && left <= right {
		if max-min == 1 {
			ans = right - left + 1
			right++
		} else {
			left++
		}
		if left == right {
			right++
		}
	}
	return ans
}
