package main

import (
	"strconv"
	"strings"
)

func toMin(time string) int {
	list := strings.Split(time, ":")
	ans, _ := strconv.Atoi(list[1])
	plus, _ := strconv.Atoi(list[0])
	ans += plus * 60
	return ans
}

func convertTime(current string, correct string) int {
	// 先计算分钟数
	currentMinutes := toMin(current)
	correctMinutes := toMin(correct)
	needle := correctMinutes - currentMinutes
	ans := 0
	nums := []int{60, 15, 5, 1}
	index := 0
	for needle > 0 {
		ans += needle / nums[index]
		needle = needle % nums[index]
		index++
	}
	return ans
}
