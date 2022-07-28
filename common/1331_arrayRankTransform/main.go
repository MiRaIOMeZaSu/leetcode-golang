package main

import "sort"

type void struct {
}

var member void

func arrayRankTransform(arr []int) []int {
	// 排序的同时去重
	set := map[int]void{}
	sortedContainer := sort.IntSlice{}

	for i := 0; i < len(arr); i++ {
		if _, ok := set[arr[i]]; !ok {
			set[arr[i]] = member
			sortedContainer = append(sortedContainer, arr[i])
		}
	}
	sort.Sort(sortedContainer)
	order := map[int]int{}
	for i := 0; i < len(sortedContainer); i++ {
		order[sortedContainer[i]] = i + 1
	}
	ans := []int{}
	for i := 0; i < len(arr); i++ {
		ans = append(ans, order[arr[i]])
	}
	return ans
}
