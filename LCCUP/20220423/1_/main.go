package main

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	ans := 0
	for i := 0; i < len(fruits); i++ {
		times := fruits[i][1] / limit
		if times*limit < fruits[i][1] {
			times++
		}
		ans += times * time[fruits[i][0]]
	}
	return ans
}
