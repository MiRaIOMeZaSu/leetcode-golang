package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findKDistantIndices(nums []int, key int, k int) []int {
	// 首先找出所有的key
	indexs := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			indexs = append(indexs, i)
		}
	}
	ans := []int{}
	for i := 0; i < len(indexs); i++ {
		left := max(0, indexs[i]-k)
		if len(ans) > 0 {
			left = max(left, ans[len(ans)-1]+1)
		}
		for j := left; j < min(indexs[i]+k+1, len(nums)); j++ {
			ans = append(ans, j)
		}
	}
	return ans
}

func main() {
	println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
	println(findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2))
}
