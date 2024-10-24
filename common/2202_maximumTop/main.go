package main

import "math"

func getMax(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func maximumTop(nums []int, k int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	if k == 0 {
		return nums[0]
	}
	if size == 1 {
		if k%2 == 0 {
			return nums[0]
		}
		return -1
	}
	if k > size {
		return getMax(nums)
	}
	end := -1
	if k < size {
		end = nums[k]
	}
	return getMax([]int{getMax(nums[:k-1]), end})
}

func main() {
	print(maximumTop([]int{35, 43, 23, 86, 23, 45, 84, 2, 18, 83, 79, 28, 54, 81, 12, 94, 14, 0, 0, 29, 94, 12, 13, 1, 48, 85, 22, 95, 24, 5, 73, 10, 96, 97, 72, 41, 52, 1, 91, 3, 20, 22, 41, 98, 70, 20, 52, 48, 91, 84, 16, 30, 27, 35, 69, 33, 67, 18, 4, 53, 86, 78, 26, 83, 13, 96, 29, 15, 34, 80, 16, 49}, 15) == 94)
	print(maximumTop([]int{2, 28, 6, 4, 3}, 1000000000) == 28)
	print(maximumTop([]int{73, 63, 62, 16, 95, 92, 93, 52, 89, 36, 75, 79, 67, 60, 42, 93, 93, 74, 94, 73, 35, 86, 96}, 59) == 96)
	print(maximumTop([]int{52, 98, 7, 10, 27, 1, 33, 17, 14, 70, 79, 41, 37, 83, 58, 69, 52, 14, 66, 7, 36, 32, 39, 69, 65, 64, 45, 90, 34, 68, 44, 51, 36, 49, 71, 54, 63, 76, 73, 76, 67, 26, 54, 76, 89, 92, 89, 69, 26, 55, 93, 89, 15, 3, 54, 91, 21, 93, 78, 29, 79, 59, 14, 80, 70, 29, 5, 80, 93, 69, 29, 22, 3, 6, 2, 36, 31, 3, 22, 96, 32, 25, 97, 82, 78, 10, 83, 46, 98, 30, 1, 93, 89}, 27) == 98)
	print(maximumTop([]int{99, 95, 68, 24, 18}, 69) == 99)
	print(maximumTop([]int{18}, 3) == -1)
	print(maximumTop([]int{2}, 1) == -1)
	print(maximumTop([]int{5, 2, 2, 4, 0, 6}, 4) == 5)
	print(maximumTop([]int{4, 6, 1, 0, 6, 2, 4}, 0) == 4)
}
