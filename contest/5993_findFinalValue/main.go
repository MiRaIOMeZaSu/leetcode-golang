package main

import "sort"

func findFinalValue(nums []int, original int) int {
	size := len(nums)
	arr := sort.IntSlice{}
	arr = append(arr, nums...)
	sort.Sort(arr)
	for i := 0; i < size; i++ {
		if arr[i] == original {
			original *= 2
		}
	}
	return original
}
