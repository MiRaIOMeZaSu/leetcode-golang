package main

import "sort"

func queue(smol, big int) int {
	if big-smol > smol {
		smol = big - smol
	}
	if smol == big {
		return 1
	}
	ans := 1
	for i := big; i > smol; i-- {
		ans *= i
	}
	ans /= 2
	return ans
}

func subArrayRanges(nums []int) int64 {
	// 重点是所有的子数组
	// 忽略长度为1的子数组
	size := len(nums)
	sorted := sort.IntSlice{}
	sorted = append(sorted, nums...)
	sort.Sort(sorted)
	ans := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			// 寻找两个区间内数字数量,进行排列组合
			midCount := j - i - 1
			count := 0
			for i := 0; i <= midCount; i++ {
				count += queue(i, midCount)
			}
			ans += count * (sorted[j] - sorted[i])
		}
	}
	return int64(ans)
}

func main() {
	subArrayRanges([]int{1, 2, 3})
}
