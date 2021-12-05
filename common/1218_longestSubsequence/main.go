package main

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func longestSubsequence(arr []int, difference int) int {
	// 使用哈希表存储序列?
	// 直接使用数组
	var table [40005]int
	ans := 1
	gap := abs(difference)
	for _, val := range arr {
		inTableVal := val + 10000 + gap
		if table[inTableVal-difference]+1 > table[inTableVal] {
			table[inTableVal] = table[inTableVal-difference] + 1
		}
		if table[inTableVal] > ans {
			ans = table[inTableVal]
		}
	}
	return ans
}
