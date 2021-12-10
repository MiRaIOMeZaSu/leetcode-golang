package main

const (
	INT_MAX = int(^uint32(0) >> 1)
)

func toLower(ch byte) byte {
	index := ch - 'a'
	if index >= 0 && index < 26 {
		// 如果index > 0
		return index
	}
	index = ch - 'A'
	if index >= 0 && index < 26 {
		// 如果index > 0
		return index
	}
	return 26
}

func shortestCompletingWord(licensePlate string, words []string) string {
	// 简单的遍历判断
	ans := ""
	pivot := INT_MAX
	var letters [26]int
	for i := 0; i < len(licensePlate); i++ {
		index := toLower(licensePlate[i])
		if index < 26 {
			letters[index]++
		}
	}
	for i := 0; i < len(words); i++ {
		var curr [26]int
		for j := 0; j < len(words[i]); j++ {
			curr[words[i][j]-'a']++
		}
		flag := true
		for j := 0; j < 26; j++ {
			if curr[j] < letters[j] {
				flag = false
				break
			}
		}
		if flag && len(words[i]) < pivot {
			ans = words[i]
			pivot = len(words[i])
		}
	}
	return ans
}

func main() {
	shortestCompletingWord("Ar16259", []string{"nature", "though", "party", "food", "any", "democratic", "building", "eat", "structure", "our"})
}
