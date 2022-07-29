package main

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func double(a int) int {
	return a * a
}

func distance(p1 []int, p2 []int) int {
	return double(p1[0]-p2[0]) + double(p1[1]-p2[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func switchPos(points []point, a, b int) {
	old := points[a]
	points[a] = points[b]
	points[b] = old
}

type point struct {
	_point   []int
	distance int
}

func findSame(points []point) bool {
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i]._point[0] == points[j]._point[0] && points[i]._point[1] == points[j]._point[1] {
				return true
			}
		}
	}
	return false
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// 从距离判断
	a := p1
	points := []point{point{p1, 0}, point{p2, 0}, point{p3, 0}, point{p4, 0}}
	if findSame(points) {
		return false
	}
	for i := 1; i < 4; i++ {
		points[i].distance = distance(a, points[i]._point)
	}
	max := max(max(points[1].distance, points[2].distance), max(points[2].distance, points[3].distance))
	if max == 0 {
		return false
	}
	for i := 1; i < 4; i++ {
		if points[i].distance == max && i != 3 {
			switchPos(points, i, 3)
		}
	}
	if points[1].distance+points[2].distance != points[3].distance {
		return false
	}
	// 如何进一步判断
	line1 := distance(points[1]._point, points[3]._point)
	line2 := distance(points[2]._point, points[3]._point)
	return line1 == line2 && line1 == points[1].distance
}

func main() {
	validSquare([]int{1, 0}, []int{0, 1}, []int{0, -1}, []int{-1, 0})
}
