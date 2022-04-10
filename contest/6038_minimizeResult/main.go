package main

import (
	"strconv"
	"strings"
)

func minimizeResult(expression string) string {
	// 遍历
	strs := strings.Split(expression, "+")
	leftNum, _ := strconv.Atoi(strs[0])
	rightNum, _ := strconv.Atoi(strs[1])
	arrLeft := [][]int{{1, leftNum}}
	arrRight := [][]int{{rightNum, 1}}
	temp := 10
	for i := 1; i < len(strs[0]); i++ {
		arrLeft = append(arrLeft, []int{leftNum / temp, leftNum % temp})
		temp *= 10
	}
	temp = 10
	for i := 1; i < len(strs[1]); i++ {
		arrRight = append(arrRight, []int{rightNum / temp, rightNum % temp})
		temp *= 10
	}
	ans := "(" + expression + ")"
	num := leftNum + rightNum
	for i := 0; i < len(arrLeft); i++ {
		for j := 0; j < len(arrRight); j++ {
			currNum := (arrLeft[i][1] + arrRight[j][0]) * arrLeft[i][0] * arrRight[j][1]
			if currNum < num {
				num = currNum
				prefix := ""
				suffix := ""
				if arrLeft[i][0] > 1 || arrLeft[i][0]*arrLeft[i][1] != leftNum {
					prefix = strconv.FormatInt(int64(arrLeft[i][0]), 10)
				}
				if arrRight[j][1] > 1 || arrRight[j][0]*arrRight[j][1] != rightNum {
					suffix = strconv.FormatInt(int64(arrRight[j][1]), 10)
				}
				ans = prefix + "(" + strconv.FormatInt(int64(arrLeft[i][1]), 10) + "+" + strconv.FormatInt(int64(arrRight[j][0]), 10) + ")" + suffix
			}
		}
	}
	return ans
}

func main() {
	minimizeResult("5+938")
	// minimizeResult("247+38")
}
