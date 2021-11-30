package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	// 循序递进
	start := 1
	curr := 0
	border, i := 10, 1
	for ; ; border, i = border*10, i+1 {
		increase := (border - start) * i
		if curr+increase > n {
			// 应该在长度为i时停下来
			break
		}
		curr += increase
		start = border
	}
	num := start + (n-curr)/i - 1
	rest := (n - curr) % i
	if rest > 0 {
		num++
		return int(strconv.Itoa(num)[rest-1] - '0')
	}
	return num % 10
}

func main() {
	fmt.Println(findNthDigit(125))
}
