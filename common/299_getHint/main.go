package main

import "strconv"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getHint(secret string, guess string) string {
	var nums1 [26]int
	var nums2 [26]int
	ans1 := 0
	ans2 := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			ans1++
		} else {
			nums1[secret[i]-'0']++
			nums2[guess[i]-'0']++
		}
	}
	for i := 0; i < 26; i++ {
		ans2 += min(nums1[i], nums2[i])
	}
	return strconv.Itoa(ans1) + "A" + strconv.Itoa(ans2) + "B"
}
