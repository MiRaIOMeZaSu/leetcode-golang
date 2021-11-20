package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLHS(nums []int) int {
	// 计算相邻的两个数
	ans := 0
	var dict map[int]int
	dict = make(map[int]int)
	for _, val := range nums {
		dict[val]++
	}
	for k, v := range dict {
		if dict[k-1] != 0 {
			ans = max(ans, dict[k-1]+v)
		}
	}
	return ans
}

func main() {
	findLHS([]int{1, 2, 3, 4, 5})
}
