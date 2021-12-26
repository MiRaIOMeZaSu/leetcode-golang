package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	list := strings.Split(text, " ")
	size := len(list)
	var ans []string
	for i := 0; i < size-2; i++ {
		if list[i] == first && list[i+1] == second {
			ans = append(ans, list[i+2])
		}
	}
	return ans
}
