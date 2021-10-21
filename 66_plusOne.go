package main

func plusOne(digits []int) []int {
	// 返回数组
	size := len(digits)
	flag := false
	for index := size - 1; index >= 0; index-- {
		next := digits[index] + 1
		if next >= 10 {
			digits[index] = next % 10
		} else {
			digits[index] = next
			flag = true
			break
		}
	}
	if !flag {
		digits = append([]int{1}, digits...)
	}
	return digits
}
