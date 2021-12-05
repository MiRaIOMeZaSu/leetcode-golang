package main

import "sort"

func targetIndices(nums []int, target int) []int {
	arr := sort.IntSlice{}
	arr = append(arr, nums...)
	sort.Sort(arr)
	var ans []int
	for i := 0; i < arr.Len(); i++ {
		if arr[i] == target {
			ans = append(ans, i)
		} else if arr[i] > target {
			return ans
		}
	}
	return ans
}
