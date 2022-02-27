package main

import "strings"

func prefixCount(words []string, pref string) int {
	size := len(words)
	ans := 0
	for i := 0; i < size; i++ {
		if strings.HasPrefix(words[i], pref) {
			ans += 1
		}
	}
	return ans
}
