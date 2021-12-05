package main

func copyArr(last [26]int) [26]int {
	var next [26]int
	for j := 0; j < 26; j++ {
		next[j] = last[j]
	}
	return next
}

func findAnagrams(s string, p string) []int {
	size := len(p)
	var ans []int
	if size > len(s) {
		return ans
	}
	var pivot [26]int
	for i := 0; i < size; i++ {
		pivot[p[i]-'a']++
	}
	var count [][26]int
	var temp [26]int
	temp[s[0]-'a'] = 1
	count = append(count, temp)
	for i := 1; i < size; i++ {
		next := copyArr(count[i-1])
		next[s[i]-'a']++
		count = append(count, next)
	}
	i := 0
	for ; i < 26; i++ {
		if pivot[i] != count[size-1][i] {
			break
		}
	}
	if i == 26 {
		ans = append(ans, 0)
	}
	// 开始计算后面的
	headIndex := 0
	currIndex := size - 1
	for j := size; j < len(s); j++ {
		next := copyArr(count[currIndex%size])
		next[s[j]-'a']++
		i = 0
		for ; i < 26; i++ {
			if next[i]-count[(currIndex+1)%size][i] != pivot[i] {
				break
			}
		}
		if i == 26 {
			ans = append(ans, headIndex+1)
		}
		count[(currIndex+1)%size] = next
		currIndex++
		headIndex++
	}
	return ans
}
