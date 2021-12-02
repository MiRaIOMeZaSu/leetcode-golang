package main

import (
	"sort"
	"strconv"
)

type element struct {
	val   int
	index int
}
type elements []element

func (s elements) Len() int {
	return len(s)
}
func (s elements) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s elements) Less(i, j int) bool {
	return s[j].val > s[i].val
}

func findRelativeRanks(score []int) []string {
	// 先排序
	// 通过map插入顺序
	medals := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	size := len(score)
	sortedContainer := []element{}
	for i := 0; i < size; i++ {
		sortedContainer = append(sortedContainer, element{score[i], i})
	}
	sort.Sort(elements(sortedContainer))
	var dict map[int]int
	dict = make(map[int]int)
	for i := 0; i < size; i++ {
		dict[sortedContainer[i].val] = sortedContainer[i].index
	}
	// 开始输出
	ans := []string{}
	for i := 0; i < size; i++ {
		rank := dict[score[i]]
		if rank >= 3 {
			ans = append(ans, strconv.Itoa(rank+1))
		} else {
			ans = append(ans, medals[rank])
		}
	}
	return ans
}

func main() {
	findRelativeRanks([]int{10, 3, 8, 9, 4})
}
