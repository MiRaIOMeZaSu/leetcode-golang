package main

func truncateSentence(s string, k int) string {
	i := 0
	for ; i < len(s) && k > 0; i++ {
		if s[i] == ' ' {
			k--
		}
	}
	if k == 0 {
		return s[:i-1]
	}
	return s[:i]
}
