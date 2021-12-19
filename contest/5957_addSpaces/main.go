package main

func addSpaces(s string, spaces []int) string {
	ans := ""
	size := len(s)
	spacesLen := len(spaces)
	index := 0
	for i := 0; i < size; i++ {
		if index < spacesLen && i == spaces[index] {
			ans += " "
			index++
		}
		ans += string(s[i])
	}
	return ans
}
