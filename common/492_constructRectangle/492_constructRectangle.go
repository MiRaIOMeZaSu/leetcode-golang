package main

const (
	INT_MAX = int(^uint(0) >> 1)
)

func constructRectangle(area int) []int {
	// 使用二份法
	left := 1
	right := area
	for left < right {
		mid := (left + right) >> 1
		pivot := mid * mid
		if mid*mid == area {
			return []int{mid, mid}
		} else if pivot < area {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	var ans []int
	ans = append(ans, []int{0, 0}...)
	for i := left; i <= area; i++ {
		if area%i == 0 {
			a := i
			b := area / a
			if a < b {
				ans[0], ans[1] = b, a
			} else {
				ans[0], ans[1] = a, b
			}
			return ans
		}
	}
	return ans
}

func main() {
	constructRectangle(122122)
}
