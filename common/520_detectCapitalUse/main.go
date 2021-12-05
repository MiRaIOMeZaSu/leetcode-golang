package main

func isUpper(ch byte) bool {
	index := ch - 'a'
	return index < 0 || index >= 26
}

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	flag1 := isUpper(word[0])
	flag2 := isUpper(word[1])
	if !flag1 && flag2 {
		return false
	}
	for i := 2; i < len(word); i++ {
		flag := isUpper(word[i])
		flag3 := flag
		if !flag2 {
			flag3 = !flag
		}
		if flag1 {
			if !flag3 {
				return false
			}
		} else if flag {
			return false
		}
	}
	return true
}

func main() {
	detectCapitalUse("USA")
}
