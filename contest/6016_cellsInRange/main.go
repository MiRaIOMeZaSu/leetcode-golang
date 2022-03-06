package main

import "strconv"

func cellsInRange(s string) []string {
	left_num := int(s[1] - '0')
	right_num := int(s[4] - '0')
	left_char := s[0]
	right_char := s[3]
	ans := []string{}
	for i := left_char; i <= right_char; i++ {
		for j := left_num; j <= right_num; j++ {
			ans = append(ans, string(i)+strconv.Itoa(j))
		}
	}
	return ans
}

func main() {
	print(cellsInRange("K1:L2"))
}
