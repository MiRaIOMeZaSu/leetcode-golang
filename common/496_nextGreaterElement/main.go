package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 使用单调栈
	size := len(nums1)
	border := len(nums2)
	var dict map[int]int = make(map[int]int)
	var stack []int
	stack = append(stack, nums2[border-1])
	dict[nums2[border-1]] = -1
	for i := border - 2; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			dict[nums2[i]] = -1
		} else {
			dict[nums2[i]] = stack[len(stack)-1]
		}
		if len(stack) == 0 || stack[len(stack)-1] > nums2[i] {
			stack = append(stack, nums2[i])
		}
	}
	for i := 0; i < size; i++ {
		nums1[i] = dict[nums1[i]]
	}
	return nums1
}

func main() {
	var num1 = []int{1, 3, 5, 2, 4}
	var num2 = []int{6, 5, 4, 3, 2, 1, 7}
	nextGreaterElement(num1, num2)
}
