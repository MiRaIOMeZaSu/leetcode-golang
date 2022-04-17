package main

import "strconv"

func digitSum(s string, k int) string {
	curr := s
	for len(curr) > k {
		next := ""
		for i := 0; i < len(curr); i += k {
			temp := 0
			for j := i; j < i+k && j < len(curr); j++ {
				num, _ := strconv.Atoi(curr[j : j+1])
				temp += num
			}
			next += strconv.Itoa(temp)
		}
		curr = next
	}
	return curr
}

func main() {
	digitSum("11111222223", 3)
}
