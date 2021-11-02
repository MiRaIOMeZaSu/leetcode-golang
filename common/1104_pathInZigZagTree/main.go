package main

import "fmt"

var ans []int
var bit []int

func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	ans = []int{}
	bit = []int{}
	// 使用数组
	// 先找到目标数字位置
	curr := 1
	count := 1
	bit = append(bit, curr)
	for i := 2; curr < label; i *= 2 {
		curr += i
		count++
		bit = append(bit, curr)
	}
	ans = append([]int{label}, ans...)
	if count%2 == 0 {
		// 此时从右到左递增
		solve((bit[count-1]-label)>>1, count-1)
	} else {
		// 此时从左到右递增
		solve((label-(bit[count-2]+1))>>1, count-1)
	}
	return ans
}

func solve(index int, level int) {
	// 使用level作为判断依据
	if level == 1 {
		ans = append([]int{1}, ans...)
		return
	}
	if level%2 == 0 {
		// 此时从右到左递增
		ans = append([]int{bit[level-1] - index}, ans...)
	} else {
		// 此时从左到右递增
		ans = append([]int{index + bit[level-2] + 1}, ans...)
	}
	solve(index>>1, level-1)
}

func main() {
	for _, val := range pathInZigZagTree(26) {
		fmt.Println(val)
	}
}
