package main

import (
	"fmt"
	"strings"
)

var blank = " "

func addSpaces(s string, spaces []int) string {
	var builder strings.Builder
	size := len(s)
	spacesLen := len(spaces)
	lastIndex := 0
	for i := 0; i < spacesLen; i++ {
		builder.WriteString(s[lastIndex:spaces[i]])
		lastIndex = spaces[i]
		builder.WriteString(blank)
	}
	builder.WriteString(s[lastIndex:size])
	return builder.String()
}

func main() {
	fmt.Print(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
}
