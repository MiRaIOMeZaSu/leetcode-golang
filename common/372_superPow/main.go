package main

import "fmt"

func compare(a []int, b []int) int {
	lenA := len(a)
	lenB := len(b)
	if lenA < lenB {
		return -1
	} else if lenA > lenB {
		return 1
	}
	// 开始逐个比较
	for i := 0; i < lenA; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mult(a []int, b []int) []int {
	ans := []int{}
	indexA := len(a) - 1
	indexB := len(b) - 1
	index := 0
	for i := indexA; i >= 0; i, index = i-1, index+1 {
		for j, currIndex := indexB, index; j >= 0; j, currIndex = j-1, currIndex+1 {
			if currIndex >= len(ans) {
				ans = append(ans, 0)
			}
			ans[currIndex] += a[indexA] * b[indexB]
		}
	}
	for i := 0; i < len(ans); i++ {
		if ans[i] >= 10 {
			if i+1 >= len(ans) {
				ans = append(ans, 0)
			}
			ans[i+1] = ans[i] / 10
			ans[i] %= 10
		}
	}
	return reverse(ans)
}

func and(a []int, b []int) []int {
	ans := []int{0}
	i, j := len(a)-1, len(b)-1
	index := 0
	for ; i >= 0 || j >= 0; i, j, index = i-1, j-1, index+1 {
		numA, numB := 0, 0
		if i >= 0 {
			numA = a[i]
		}
		if j >= 0 {
			numB = b[j]
		}
		ans[index] += numA + numB
		if ans[index] >= 10 {
			ans[index] %= 10
			ans = append(ans, 1)
		} else {
			ans = append(ans, 0)
		}
	}
	if ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	return reverse(ans)
}

func sub(a []int, b []int) []int {
	ans := []int{}
	// 用a减b
	i := 0
	// 开始减
	indexA := len(a) - 1
	indexB := len(b) - 1
	for i, j := indexB, indexA; i >= 0; i, j = i-1, j-1 {
		a[indexA] -= b[indexB]
		if a[indexA] < 0 {
			k := j - 1
			for ; j >= 0 && a[k] == 0; k-- {
				a[k] = 9
			}
			a[k] -= 1
			a[indexA] += 10
		}
	}
	// 缩减0
	for ; i < len(a) && a[i] == 0; i++ {
	}
	for ; i < len(a); i++ {
		ans = append(ans, a[i])
	}
	return ans
}

func reverse(a []int) []int {
	left := 0
	right := len(a) - 1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left++
		right--
	}
	return a
}

func solve(next []int, rest []int, pivot []int) []int {
	flag := compare(next, rest)
	if flag == -1 {
		// rest 更大
		rest := solve(mult(pivot, next), rest, pivot)
		flag = compare(next, rest)
		if flag == 0 {
			return []int{}
		} else if flag == 1 {
			return sub(rest, next)
		}
	} else if flag == 0 {
		return []int{}
	}
	// 若next更大则不作处理
	return rest
}

func superPow(a int, b []int) int {
	// 更复杂的数组乘法
	// a整体的乘会出现溢出
	pivot := []int{}
	for i := a; i > 0; i /= 10 {
		pivot = append(pivot, a%10)
	}
	pivot = reverse(pivot)
	curr := []int{}
	curr = append(curr, pivot...)
	// 使用回溯法?
	ans := solve(pivot, b, pivot)
	if len(ans) == 0 {
		return 0
	}
	return ans[0]
}

func main() {
	// 开始测试
	fmt.Println(and([]int{1, 2, 3, 4, 5}, []int{9, 9, 9, 9, 1, 2, 4, 5, 1}))
	fmt.Println(sub([]int{9, 9, 9, 9, 1, 2, 4, 5, 1}, []int{1, 2, 3, 4, 5}))
	fmt.Println(mult([]int{9, 9, 9, 9, 1, 2, 4, 5, 1}, []int{1, 2, 3, 4, 5}))
	fmt.Println(superPow(3, []int{8}))
}
