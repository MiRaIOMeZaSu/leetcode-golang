package main

func maxDistance(colors []int) int {
	// 求出每种颜色的位置
	size := len(colors)
	var houseColor map[int]([]int)
	houseColor = make(map[int][]int)
	for index, val := range colors {
		if houseColor[val] == nil {
			houseColor[val] = []int{}
		}
		houseColor[val] = append(houseColor[val], index)
	}
	ans := 0
	for color, arr := range houseColor {
		for i := 0; i < size; i++ {

		}
	}
	return ans
}
