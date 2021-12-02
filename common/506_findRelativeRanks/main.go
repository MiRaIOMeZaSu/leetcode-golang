package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	// 先排序
	// 通过map插入顺序
	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	size := len(score)
	sortedContainer := sort.IntSlice{}
	for i := 0; i < size; i++ {
		sortedContainer = append(sortedContainer, score[i])
	}
	sort.Sort(sortedContainer)
	var dict map[int]int
	dict = make(map[int]int)
	for i := 0; i < size; i++ {
		dict[sortedContainer[i]] = size - i
	}
	// 开始输出
	ans := []string{}
	for i := 0; i < size; i++ {
		rank := dict[score[i]]
		if rank > 3 {
			ans = append(ans, strconv.Itoa(rank))
		} else {
			ans = append(ans, medals[rank-1])
		}
	}
	return ans
}

func main() {
	findRelativeRanks([]int{10, 3, 8, 9, 4})
}
