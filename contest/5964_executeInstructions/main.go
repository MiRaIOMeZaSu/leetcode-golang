package main

func executeInstructions(n int, startPos []int, s string) []int {
	ans := []int{}
	size := len(s)
	// 存储在某个位置时可以执行的数量?
	currPos := []int{}

	for i := 0; i < size; i++ {
		currPos = []int{}
		currPos = append(currPos, startPos...)
		temp := 0
		for j := i; j < size; j++ {
			if s[j] == 'U' {
				if currPos[0] > 0 {
					currPos[0]--
					temp++
				} else {
					break
				}
			} else if s[j] == 'D' {
				if currPos[0] != n-1 {
					currPos[0]++
					temp++
				} else {
					break
				}
			} else if s[j] == 'L' {
				if currPos[1] > 0 {
					currPos[1]--
					temp++
				} else {
					break
				}
			} else if s[j] == 'R' {
				if currPos[1] != n-1 {
					currPos[1]++
					temp++
				} else {
					break
				}
			}
		}
		ans = append(ans, temp)
	}
	return ans
}

func main() {
	executeInstructions(3, []int{0, 1}, "RRDDLU")
}
