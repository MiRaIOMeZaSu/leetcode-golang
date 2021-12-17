package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	curr := numBottles
	for curr >= numExchange {
		next := curr / numExchange
		ans += next
		curr = curr%numExchange + next
	}
	return ans
}

func main() {
	fmt.Print(numWaterBottles(9, 3))
}
