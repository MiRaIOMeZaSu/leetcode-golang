package main

func maxDistance(colors []int) int {
	// 求出每种颜色的位置
	size := len(colors)
	for i := size - 1; i >= 1; i-- {
		for j := 0; j+i < size; j++ {
			if colors[j] != colors[j+i] {
				return i
			}
		}
	}
	return 1
}
