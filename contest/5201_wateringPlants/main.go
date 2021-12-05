package main

import "fmt"

func wateringPlants(plants []int, capacity int) int {
	// 使用遍历
	count := len(plants)
	ans := 0
	curr := capacity
	for i := 0; i < count; i++ {
		if curr >= plants[i] {
			curr -= plants[i]
		} else {
			curr = capacity - plants[i]
			ans += i * 2
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Print(wateringPlants([]int{2, 2, 3, 3}, 5))
}
