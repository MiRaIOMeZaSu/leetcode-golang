package main

func getAverages(nums []int, k int) []int {
	// 使用前缀和
	var ans []int
	var prex []int
	size := len(nums)
	prex = append(prex, 0)
	for i := 0; i < size; i++ {
		prex = append(prex, prex[i]+nums[i])
	}
	for i := 0; i < size; i++ {
		if i-k >= 0 && i+k < size {
			ans = append(ans, (prex[i+k+1]-prex[i-k])/(k+k+1))
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}

func main() {
	getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3)
}
