package main

import "math"

func getMax(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func maximumTop(nums []int, k int) int {
	// x: in hand, y: remain times
	var graph [][]int
	size := len(nums)
	for i := 0; i <= size; i++ {
		graph = append(graph, []int{})
		for j := 0; j <= k; j++ {
			graph[i] = append(graph[i], -1)
		}
	}
	max := -1
	if size > k {
		max = nums[k]
	}
	graph[1][k-1] = nums[0]
	for i := 2; i <= size && k-i-1 >= 0; i++ {
		graph[i][k-i] = getMax([]int{nums[i-1], graph[i-1][k-i+1]})
	}

	for i := k - 2; i >= 0; i-- {
		for j := 2; j < size; j++ {
			graph[j][i] = getMax([]int{
				graph[j][i],      // origin
				graph[j+1][i+1],  // drop one
				graph[j-1][i+1]}) // pick one
		}
	}
	for i := 1; i <= size; i++ {
		if graph[i][1] > max {
			max = graph[i][1]
		}
	}
	return max
}

func main() {
	print(maximumTop([]int{35, 43, 23, 86, 23, 45, 84, 2, 18, 83, 79, 28, 54, 81, 12, 94, 14, 0, 0, 29, 94, 12, 13, 1, 48, 85, 22, 95, 24, 5, 73, 10, 96, 97, 72, 41, 52, 1, 91, 3, 20, 22, 41, 98, 70, 20, 52, 48, 91, 84, 16, 30, 27, 35, 69, 33, 67, 18, 4, 53, 86, 78, 26, 83, 13, 96, 29, 15, 34, 80, 16, 49}, 15))
}
