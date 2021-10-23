package _69_majorityElement

func majorityElement(nums []int) int {
	var a int
	vote1 := 0
	for i := 0; i < len(nums); i++ {
		if vote1 == 0 {
			a = nums[i]
			vote1++
		} else if a == nums[i] {
			vote1++
		} else {
			vote1--
		}
	}
	// 此时a为最大数量的值
	return a
}
