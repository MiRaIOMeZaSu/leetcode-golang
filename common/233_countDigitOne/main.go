package main

import "fmt"

func countDigitOne(n int) int {
	// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
	// 对于当前位数,该值容易算出
	// 困难点是中间位数为1的处理方法?
	// 前缀可以为0,后缀同样可以为0
	// 当前值大于1时情况会变得复杂,要考虑后缀的多种情况?
	// 当前值为0时更复杂
	if n == 0 {
		return 0
	}
	ans := n / 10
	if n%10 > 0 {
		ans += 1
	}
	// 开始计算第二位的
	pivot := 10
	for ; n >= pivot; pivot *= 10 {
		right := n % pivot
		left := n / (pivot * 10)
		curr := ((n % (pivot * 10)) - right) / pivot
		if curr == 0 {
			ans += left * pivot
		} else if curr == 1 {
			ans += (right + 1) + left*pivot
		} else {
			ans += (left + 1) * pivot
		}
	}
	return ans
}

func main() {
	fmt.Println(countDigitOne(110))
}
