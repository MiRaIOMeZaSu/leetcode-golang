package main

import (
	"fmt"
	"strconv"
)

func getCount(n int) (int, [10]int) {
	var numCount [10]int
	count := 0
	curr := n
	for ; curr != 0; count++ {
		numCount[curr%10]++
		curr /= 10
	}
	return count, numCount
}

func reorderedPowerOf2(n int) bool {
	// 只需要计算数字的数量是否匹配

	count, numCount := getCount(n)
next:
	for i := 1; ; i *= 2 {
		str := strconv.Itoa(i)
		length := len(str)
		if length == count {
			_, thisCount := getCount(i)
			for j := 0; j < 10; j++ {
				if thisCount[j] != numCount[j] {
					continue next
				}
			}
			return true
		} else if len(str) > count {
			return false
		}
	}
}

func main() {
	fmt.Print(reorderedPowerOf2(3))
}
