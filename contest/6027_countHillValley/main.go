package main

func satisfy(nums []int, index int) bool {
	left := index - 1
	for ; left >= 0 && nums[left] == nums[index]; left-- {
	}
	right := index + 1
	for ; right < len(nums) && nums[right] == nums[index]; right++ {
	}
	if left < 0 || right >= len(nums) {
		return false
	}
	if nums[left] > nums[index] && nums[right] > nums[index] {
		return true
	}
	if nums[left] < nums[index] && nums[right] < nums[index] {
		return true
	}
	return false
}

func countHillValley(nums []int) int {
	// 左最大和右最小?
	ans := 0
	lastSatisfy := false
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == nums[i-1] && lastSatisfy {
			continue
		}
		if satisfy(nums, i) {
			lastSatisfy = true
			ans += 1
		}
	}
	return ans
}

func main() {
	print(countHillValley([]int{5, 7, 7, 1, 7}))
	print(countHillValley([]int{2, 4, 1, 1, 6, 5}))
}
