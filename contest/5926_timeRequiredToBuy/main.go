package main

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func timeRequiredToBuy(tickets []int, k int) int {
	// 使用判断
	ans := 0
	pivot := tickets[k] - 1
	if pivot > 0 {
		for index, val := range tickets {
			temp := min(pivot, val)
			tickets[index] -= temp
			ans += temp
		}
	}
	for i := 0; i <= k; i++ {
		if tickets[i] > 0 {
			ans += 1
		}
	}
	return ans
}
