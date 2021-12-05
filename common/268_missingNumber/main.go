package main

import "fmt"

func missingNumber(nums []int) int {
	// 因为只有一个数，使用异或即可
	size := len(nums)
	curr := 0
	for index, val := range nums {
		curr ^= val
		curr ^= index
	}
	curr ^= size
	return curr
}

func main() {
	fmt.Print(missingNumber([]int{3, 0, 1}))
}
