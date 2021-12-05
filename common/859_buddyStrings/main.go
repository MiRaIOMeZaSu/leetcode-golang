package main

func buddyStrings(s string, goal string) bool {
	// 有两个相同字母,或有两个不同字母
	// 其他的必须对应
	size := len(s)
	if size != len(goal) {
		return false
	}
	diff := 0
	var alp1 [26]int
	var alp2 [26]int
	for i := 0; i < size && diff <= 2; i++ {
		alp1[s[i]-'a'] += 1
		alp2[goal[i]-'a'] += 1
		if s[i] != goal[i] {
			diff++
		}
	}
	if diff > 2 || diff%2 == 1 {
		return false
	} else if diff == 2 {
		for i := 0; i < 26; i++ {
			if alp1[i] != alp2[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < 26; i++ {
		if alp1[i] >= 2 {
			return true
		}
	}
	return false
}
