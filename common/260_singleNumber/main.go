package main

import "sort"

func singleNumber(nums []int) []int {
	var sortedArr = sort.IntSlice{}
	sortedArr = append(sortedArr, nums...)
	sort.Sort(sortedArr)
	var ans []int
	i := 0
	last := sortedArr[i]
	i++
	for i < len(sortedArr) && len(ans) < 2 {
		// 从前到后
		if sortedArr[i] == last {
			if i+1 >= len(sortedArr) {
				break
			}
			last = sortedArr[i+1]
			i += 2
		} else {
			ans = append(ans, last)
			last = sortedArr[i]
			i++
		}
	}
	if len(ans) < 2 {
		ans = append(ans, last)
	}
	return ans
}

func main() {
	singleNumber([]int{1, 0})
}
