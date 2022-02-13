package main

func countOperations(num1 int, num2 int) int {
	ans := 0
	for num1 != 0 && num2 != 0 {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		ans += 1
	}
	return ans
}

func main() {
	countOperations(2, 3)
}
