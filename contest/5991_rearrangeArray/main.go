package main

func rearrangeArray(nums []int) []int {
	// 以正整数开头
	size := len(nums)
	negtive := []int{}
	postive := []int{}
	for i := 0; i < size; i++ {
		if nums[i] < 0 {
			negtive = append(negtive, nums[i])
		}
	}
	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			postive = append(postive, nums[i])
		}
	}
	ans := []int{}
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			ans = append(ans, postive[i/2])
		} else {
			ans = append(ans, negtive[i/2])
		}
	}
	return ans
}
