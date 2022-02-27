package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minSteps(s string, t string) int {
	// 简单地统计相同字符串即可
	var s1 [26]int
	var s2 [26]int
	ans := 0
	s1_size := len(s)
	s2_size := len(t)
	pviot := 'a'
	for i := 0; i < s1_size; i++ {
		s1[s[i]-byte(pviot)] += 1
	}
	for i := 0; i < s2_size; i++ {
		s2[t[i]-byte(pviot)] += 1
	}
	for i := 0; i < 26; i++ {
		ans += abs(s1[i] - s2[i])
	}
	return ans
}

func main() {
	minSteps("leetcode", "coats")
}
