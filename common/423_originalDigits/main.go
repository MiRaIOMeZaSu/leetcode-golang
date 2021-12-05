package main

import (
	"fmt"
	"strconv"
	"strings"
)

func originalDigits(s string) string {
	// 多重背包
	// 超级长!
	// 特殊组合数量是固定的!
	// w: two g: eight z: zero x: six u: four r: three o: one f: five v: seven n:nine
	var letterCount [26]int
	var numsCount [10]int
	// 开始统计
	for i := 0; i < len(s); i++ {
		letterCount[s[i]-'a']++
	}
	var words = [10][3]string{
		{"w", "two", "2"},
		{"g", "eight", "8"},
		{"z", "zero", "0"},
		{"x", "six", "6"},
		{"u", "four", "4"},
		{"r", "three", "3"},
		{"o", "one", "1"},
		{"f", "five", "5"},
		{"v", "seven", "7"},
		{"i", "nine", "9"}}
	for i := 0; i < len(words); i++ {
		toMinus := letterCount[words[i][0][0]-'a']
		if toMinus > 0 {
			num := words[i][2][0] - '0'
			for j := 0; j < len(words[i][1]); j++ {
				letterCount[words[i][1][j]-'a'] -= toMinus
			}
			numsCount[num] += toMinus
		}
	}
	ans := ""
	for i := 0; i < len(numsCount); i++ {
		ans += strings.Repeat(strconv.Itoa(i), numsCount[i])
	}
	return ans
}

func main() {
	fmt.Print(originalDigits("nnei"))
}
