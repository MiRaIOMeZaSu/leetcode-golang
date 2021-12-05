package main

import (
	"sort"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func largestSumAfterKNegations(nums []int, k int) int {
	// 选择所有最小的数字取反即可
	size := len(nums)
	sortedContainer := sort.IntSlice{}
	sortedContainer = append(sortedContainer, nums...)
	sort.Sort(sortedContainer)
	ans := 0
	// 可以选择相同的重复多次
	// 先寻找负数
	currK := k
	negativeCount := 0
	for i := 0; i < size && currK > 0 && sortedContainer[i] < 0; i, currK = i+1, currK-1 {
		ans -= sortedContainer[i]
		negativeCount++
	}
	// 剩下的
	currK %= 2
	if negativeCount < size {
		// 还剩下正数
		if currK == 1 {
			// 如何减去?
			if negativeCount-1 < 0 || abs(sortedContainer[negativeCount-1]) > abs(sortedContainer[negativeCount]) {
				ans -= sortedContainer[negativeCount]
			} else {
				ans += sortedContainer[negativeCount-1] * 2
				ans += sortedContainer[negativeCount]
			}
		} else {
			ans += sortedContainer[negativeCount]
		}
	} else if currK == 1 {
		ans += sortedContainer[negativeCount-1] * 2
	}
	for i := negativeCount + 1; i < size; i++ {
		ans += sortedContainer[i]
	}
	return ans
}

func main() {
	largestSumAfterKNegations([]int{-8, 3, -5, -3, -5, -2}, 6)
}
