package main

import "sort"

const (
	INT_MAX = int(^uint32(0) >> 1)
)

func nextGreaterElement(n int) int {
	// 使用排序
	var nums []int
	curr := n
	for curr > 0 {
		nums = append(nums, curr%10)
		curr /= 10
	}
	arr := sort.IntSlice{}
	arr = append(arr, nums...)
	sort.Sort(arr)
	var ans int
	ans = 0
	for i := len(arr) - 1; i >= 0; i-- {
		ans *= 10
		ans += arr[i]
	}
	if ans == n || ans > INT_MAX {
		return -1
	}
	return ans
}

func main() {
	nextGreaterElement(21)
}
