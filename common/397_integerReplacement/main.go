package main

import "fmt"

const (
	INT_MAX = int(^uint32(0) >> 1)
)

var mem map[int]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func integerReplacement(n int) int {
	if mem == nil {
		mem = make(map[int]int)
	}
	// 动态规划?
	if val, ok := mem[n]; ok {
		return val
	}
	mem[n] = INT_MAX
	curr := n
	ans := 0
	for curr != 1 {
		if curr%2 == 0 {
			curr /= 2
		} else {
			mem[n] = min(mem[n], ans+2+integerReplacement((curr-1)/2))
			mem[n] = min(mem[n], ans+2+integerReplacement((curr+1)/2))
			return mem[n]
		}
		ans++
	}
	mem[n] = min(mem[n], ans)
	return mem[n]
}

func main() {
	fmt.Println(INT_MAX)
	fmt.Println(integerReplacement(6))
}
