package main

func countEven(num int) int {
	// 只要首字母不为零,都可以
	// 求的是数目
	// 遍历即可
	ans := 0
	for i := 1; i <= num; i++ {
		temp := 0
		curr_i := i
		for curr_i > 0 {
			temp += curr_i % 10
			curr_i /= 10
		}
		if temp%2 == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	countEven(4)
}
