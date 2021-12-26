package main

func executeInstructions(n int, startPos []int, s string) []int {
	ans := []int{}
	size := len(s)
	for i := 0; i < size; i++ {
		ans = append(ans, 0)
	}
	// 存储在某个位置时可以执行的数量?
	// var dict map[int]string
	currPos := []int{}

	// dict = make(map[int]string)
	for i := size - 1; i >= 0; i-- {
		currPos = []int{}
		currPos = append(currPos, startPos...)
		temp := 0
		for j := i; j < size; j++ {
			if s[j] == 'U' && currPos[1] > 0 {
				currPos[1]--
				temp++
			} else if s[j] == 'D' && currPos[1] != n-1 {
				currPos[1]++
				temp++
			} else if s[j] == 'L' && currPos[0] > 0 {
				currPos[0]--
				temp++
			} else if s[j] == 'R' && currPos[0] != n-1 {
				currPos[0]++
				temp++
			}
			if currPos[0] == startPos[0] && currPos[1] == startPos[1] {
				ans[i] = temp + ans[j]
				break
			}
		}
		if ans[i] == 0 {
			ans[i] = temp
		}
	}
	return ans
}
