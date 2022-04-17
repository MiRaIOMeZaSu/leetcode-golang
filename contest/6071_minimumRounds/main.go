package main

func minimumRounds(tasks []int) int {
	// 统计数量
	var table map[int]int
	table = make(map[int]int)
	for i := 0; i < len(tasks); i++ {
		table[tasks[i]]++
	}
	ans := 0
	for _, v := range table {
		if v == 1 {
			return -1
		}
		if v%3 != 0 {
			ans += v/3 + 1
		} else {
			ans += v / 3
		}
	}
	return ans
}

func main() {
	minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4})
}
