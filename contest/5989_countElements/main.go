package main

import "sort"

func countElements(nums []int) int {
	arr := sort.IntSlice{}
	arr = append(arr, nums...)
	sort.Sort(arr)
	size := len(nums)
	ans := 0
	max := arr[size-1]
	min := arr[0]
	for i := 1; i < size-1; i++ {
		if arr[i] > min && arr[i] < max {
			ans += 1
		}
	}
	return ans
}

func main() {
	countElements([]int{-3, 3, 3, 90})
}
