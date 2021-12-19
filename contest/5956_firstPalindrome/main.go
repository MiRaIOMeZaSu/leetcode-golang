package main

func firstPalindrome(words []string) string {
	size := len(words)
	for i := 0; i < size; i++ {
		left := 0
		right := len(words[i]) - 1
		isPalindrome := true
		for left < right {
			if words[i][left] != words[i][right] {
				isPalindrome = false
				break
			}
			left++
			right--
		}
		if isPalindrome {
			return words[i]
		}
	}
	return ""
}
