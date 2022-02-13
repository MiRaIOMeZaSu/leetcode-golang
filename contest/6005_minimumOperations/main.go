package main

import "sort"

const (
	INT_MAX = int(^uint32(0) >> 1)
)

type Number struct {
	Val   int
	Count int
}

type Numbers []*Number

func (numbers Numbers) Len() int {
	return len(numbers)
}

func (numbers Numbers) Less(i, j int) bool {
	return numbers[i].Count > numbers[j].Count
}

func (numbers Numbers) Swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func minimumOperations(nums []int) int {
	// 统计数量
	// 只有10 * 10 种组合
	if len(nums) == 0 {
		return 0
	}
	var first map[int]int
	first = make(map[int]int)
	size := len(nums)
	for i := 0; i < size; i += 2 {
		first[nums[i]]++
	}
	var numNodesFirst Numbers
	for k, v := range first {
		numNodesFirst = append(numNodesFirst, &Number{k, v})
	}
	var second map[int]int
	second = make(map[int]int)
	for i := 1; i < size; i += 2 {
		second[nums[i]]++
	}
	var numNodesSecond Numbers
	for k, v := range second {
		numNodesSecond = append(numNodesSecond, &Number{k, v})
	}
	sort.Sort(numNodesFirst)
	sort.Sort(numNodesSecond)
	ans := INT_MAX
	// 开始组合
	firstSize := len(numNodesFirst)
	secondSize := len(numNodesSecond)
	for i := 0; i < firstSize; i++ {
		j := 0
		for ; j < secondSize; j++ {
			if numNodesFirst[i].Val != numNodesSecond[j].Val {
				ans = min(ans, size-numNodesFirst[i].Count-numNodesSecond[j].Count)
				break
			}
		}
		if j == secondSize {
			ans = min(ans, size-numNodesFirst[i].Count)
		}
	}
	return ans
}

func main() {
	minimumOperations([]int{2, 2})
}
