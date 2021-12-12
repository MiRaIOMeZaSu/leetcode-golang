package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func subArrayRanges(nums []int) int64 {
	// 重点是所有的子数组
	// 忽略长度为1的子数组
	size := len(nums)
	ans := 0
	for i := 0; i < size; i++ {
		minNum := nums[i]
		maxNum := nums[i]
		for j := i + 1; j < size; j++ {
			minNum = min(minNum, nums[j])
			maxNum = max(maxNum, nums[j])
			// 寻找两个区间内数字数量,进行排列组合
			ans += maxNum - minNum
		}
	}
	return int64(ans)
}

func main() {
	subArrayRanges([]int{1, 2, 3})
}
