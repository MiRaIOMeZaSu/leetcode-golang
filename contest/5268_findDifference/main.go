package main

type void struct{}

var member void

func findDifference(nums1 []int, nums2 []int) [][]int {
	ans := [][]int{}
	ans = append(ans, []int{})
	ans = append(ans, []int{})
	var map1 map[int]void
	map1 = make(map[int]void)
	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]] = member
	}
	var map2 map[int]void
	map2 = make(map[int]void)
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]] = member
		if _, ok := map1[nums2[i]]; !ok {
			map1[nums2[i]] = member
			ans[1] = append(ans[1], (nums2[i]))
		}
	}
	for i := 0; i < len(nums1); i++ {
		if _, ok := map2[nums1[i]]; !ok {
			map2[nums1[i]] = member
			ans[0] = append(ans[0], (nums1[i]))
		}
	}
	return ans
}
