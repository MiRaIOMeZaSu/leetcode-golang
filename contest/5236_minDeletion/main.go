package main

func minDeletion(nums []int) int {
	// 必须是偶数
	// 则要必要时删除最后
	ans := 0
	right := 1
	for i := 0; right < len(nums); {
		if nums[i] == nums[right] {
			// 此时移动
			right++
			ans++
		} else {
			i = right + 1
			right = i + 1
		}
	}
	// 删除的数目
	if (len(nums)-ans)%2 != 0 {
		ans++
	}
	return ans
}

func main() {
	minDeletion([]int{2, 6, 2, 5, 8, 9, 7, 2, 2, 5, 6, 2, 2, 0, 6, 8, 7, 3, 9, 2, 1, 1, 3, 2, 6, 2, 4, 6, 5, 8, 4, 8, 7, 0, 4, 8, 7, 8, 4, 1, 1, 4, 0, 1, 5, 7, 7, 5, 9, 7, 5, 5, 8, 6, 4, 3, 6, 5, 1, 6, 7, 6, 9, 9, 6, 8, 6, 0, 9, 5, 6, 7, 6, 9, 5, 5, 7, 3, 0, 0, 5, 5, 4, 8, 3, 9, 3, 4, 1, 7, 9, 3, 1, 8, 8, 9, 1, 6, 0, 0})
}
