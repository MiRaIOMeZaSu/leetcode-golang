package nextgreaterelement

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 使用单调栈
	size := len(nums1)
	border := len(nums2)
	var stack []int
	stack = append(stack, nums2[border-1])
	for i := border - 2; i >= size; i-- {
		for len(stack) != 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:1]
		}
		if len(stack) == 0 || stack[len(stack)-1] > nums2[i] {
			stack = append(stack, nums2[i])
		}
	}
	for i := size - 1; i >= 0; i-- {
		curr := nums1[i]
		for curr >= stack[len(stack)-1] {

		}
	}

	return nums1
}
