package main

func maxScoreIndices(nums []int) []int {
	size := len(nums)
	// 前缀和
	prexZero := []int{0}
	prexOne := []int{0}
	for i := 0; i < size; i++ {
		prexOne = append(prexOne, prexOne[i])
		prexZero = append(prexZero, prexZero[i])
		if nums[i] == 0 {
			prexZero[i+1] += 1
		} else {
			prexOne[i+1] += 1
		}
	}
	// 开始遍历
	max := 0
	ans := []int{}
	for i := 0; i <= size; i++ {
		curr := prexZero[i] + (prexOne[size] - prexOne[i])
		if curr > max {
			ans = []int{}
			ans = append(ans, i)
			max = curr
		} else if curr == max {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	maxScoreIndices([]int{0, 0, 1, 0})
}
