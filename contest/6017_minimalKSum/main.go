package main

import "sort"

func minimalKSum(nums []int, k int) int64 {
	arr := sort.IntSlice{}
	arr = append(arr, nums...)
	sort.Sort(arr)
	var ans int64
	ans = 0
	curr := 1
	index := 0
	size := len(arr)
	for k > 0 && index < size {
		if arr[index] == curr {
			for index < size && arr[index] == curr {
				index++
			}
			curr++
		} else if arr[index] > curr {
			ans += int64(curr)
			curr++
			k--
		}
	}
	for k > 0 {
		ans += int64(curr)
		curr++
		k--
	}
	return ans
}

func main() {
	// print(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
	// print(minimalKSum([]int{5, 6}, 6))
	// print(minimalKSum([]int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52, 19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84}, 35))
	print(minimalKSum([]int{1000000000}, 1000000000))
}
