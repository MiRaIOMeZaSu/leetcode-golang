package main

import (
	"fmt"
	"strconv"
)

type void struct {
}

var member void

type Number struct {
	negtivate int
	up        int
	down      int
}

const (
	PLUS      = '+'
	NEGTIVATE = '-'
	MIN_NUM   = '0'
	MAX_NUM   = '9'
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func greatestCommonDivisor(a, b int) int {
	maxNum := max(a, b)
	minNum := min(a, b)
	if maxNum%minNum == 0 {
		return minNum
	}
	ans := 1
	for i := 2; i <= minNum && minNum > 1; i++ {
		for minNum%i == 0 && maxNum%i == 0 {
			ans *= i
			minNum /= i
			maxNum /= i
		}
	}
	return ans
}

func _leastCommmonMultiple(a, b int) int {
	divisor := greatestCommonDivisor(a, b)
	return a * b / divisor
}

func leastCommmonMultiple(list []int) int {
	if len(list) == 1 {
		return list[0]
	}
	length := len(list)
	nums := []int{}
	if length%2 == 0 {
		for i := 0; i+1 < len(list); i++ {
			nums = append(nums, _leastCommmonMultiple(list[i], list[i+1]))
		}
		return leastCommmonMultiple(nums)
	}
	return _leastCommmonMultiple(list[0], leastCommmonMultiple(list[1:]))
}

func fractionAddition(expression string) string {
	chars := []rune(expression)
	nums := []Number{}
	set := map[rune]void{PLUS: member, NEGTIVATE: member}
	for i := 0; i < len(chars); {
		num := Number{1, 0, 0}
		if _, ok := set[chars[i]]; ok {
			// 此时可以分割
			if chars[i] != PLUS {
				num.negtivate = -1
			}
			i++
		}
		for chars[i] <= MAX_NUM && chars[i] >= MIN_NUM {
			num.up *= 10
			num.up += int(chars[i] - MIN_NUM)
			i++
		}
		i++
		for i < len(chars) && chars[i] <= MAX_NUM && chars[i] >= MIN_NUM {
			num.down *= 10
			num.down += int(chars[i] - MIN_NUM)
			i++
		}
		nums = append(nums, num)
	}
	// 取出来
	// 最小公倍数
	downNums := []int{}
	for i := 0; i < len(nums); i++ {
		downNums = append(downNums, nums[i].down)
	}
	downCommon := leastCommmonMultiple(downNums)
	up := 0
	ans := ""
	for i := 0; i < len(nums); i++ {
		magnification := downCommon / nums[i].down
		up += nums[i].negtivate * nums[i].up * magnification
	}
	if up < 0 {
		ans += string(NEGTIVATE)
	} else if up == 0 {
		return "0/1"
	}
	div := greatestCommonDivisor(abs(up), downCommon)
	return ans + strconv.Itoa(abs(up)/div) + "/" + strconv.Itoa(downCommon/div)
}

func main() {
	fmt.Println(fractionAddition("-1/4-4/5-1/4"))
}
