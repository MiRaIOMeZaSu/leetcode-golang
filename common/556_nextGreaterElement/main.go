package main

import "sort"

const (
	INT_MAX = int(^uint32(0) >> 1)
)

func nextGreaterElement(n int) int {
	// 使用排序
	// 只需要使得靠后的某位的值比原来大
	nums := sort.IntSlice{}
	curr := n
	var ans int
	ans = -1
	max := 0
	index := 0
	currIndex := 0
	for curr > 0 {
		num := curr % 10
		if max > num {
			// 可以开始替换
			min := max
			for i := 0; i < len(nums); i++ {
				if nums[i] > num && nums[i] < min {
					min = nums[i]
					index = i
				}
			}
			nums = append(nums[:index], nums[index+1:]...)
			nums = append(nums, num)
			ans = curr / 10
			ans *= 10
			ans += min
			sort.Sort(nums)
			for i := 0; i < len(nums); i++ {
				ans *= 10
				ans += nums[i]
			}
			break
		} else {
			max = num
			index = currIndex
		}
		nums = append(nums, num)
		curr /= 10
		currIndex++
	}

	if ans == n || ans > INT_MAX {
		return -1
	}
	return ans
}

func main() {
	nextGreaterElement(21)
}
