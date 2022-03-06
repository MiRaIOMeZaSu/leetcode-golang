package main

import "sort"

func min(a int, b int) int64 {
	if a > b {
		return int64(b)
	}
	return int64(a)
}

func minimalKSum(nums []int, k int) int64 {
	arr := sort.IntSlice{}
	arr = append(arr, nums...)
	sort.Sort(arr)
	var ans int64
	ans = 0
	left_index := 0
	size := len(arr)
	if arr[0] > 1 {
		gap := min(k, arr[0]-1)
		// 从1到gap的所有值
		k -= int(gap)
		ans += ((1 + gap) * gap) / 2
	}
	for k > 0 && left_index+1 <= size-1 {
		// 直接对区间求值
		gap := min(k, arr[left_index+1]-arr[left_index]-1)
		// 从arr[left_index] + 1到arr[left_index] + 1 + gap - 1
		if gap > 0 {
			ans += gap * (2*int64(arr[left_index]) + 1 + gap) / 2
			k -= int(gap)
		}
		left_index++
	}
	if k > 0 {
		gap := int64(k)
		ans += gap * (2*int64(arr[size-1]) + 1 + gap) / 2
	}
	return ans
}

func main() {
	// print(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
	// print(minimalKSum([]int{5, 6}, 6))
	print(minimalKSum([]int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52, 19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84}, 35))
	// print(minimalKSum([]int{1000000000}, 1000000000))
}
