package main

const (
	INT_MAX = int(^uint((0)) >> 1)
)

func minMoves(nums []int) int {
	min := INT_MAX
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	ans := 0
	for i := 0; i < size; i++ {
		ans += nums[i] - min
	}
	return ans
}
