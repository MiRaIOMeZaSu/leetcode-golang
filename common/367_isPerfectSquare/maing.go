package main

func isPerfectSquare(num int) bool {
	// 使用二分法
	left := 1
	right := num
	for left <= right {
		mid := (left + right) >> 1
		pivot := mid * mid
		if pivot < num {
			left = mid + 1
		} else if pivot > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	isPerfectSquare(1)
}
