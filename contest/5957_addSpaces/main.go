package main

import "strings"

func addSpaces(s string, spaces []int) string {
	var builder strings.Builder
	size := len(s)
	spacesLen := len(spaces)
	index := 0
	for i := 0; i < size; i++ {
		if index < spacesLen && i == spaces[index] {
			builder.WriteString(" ")
			index++
		}
		builder.WriteString(s[i:i])
	}
	return builder.String()
}

func main() {
	addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15})
}
