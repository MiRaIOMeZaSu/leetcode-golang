package main

func integerReplacement(n int) int {
	// 动态规划?
	// 只使用减一和除二操作
	curr := n
	ans := 0
	for curr != 1 {
		if curr%2 == 0 {
			curr /= 2
		} else {
			curr--
		}
		ans++
	}
	return ans
}
