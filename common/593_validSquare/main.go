package main

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// 先找平行的两点, 得出边
	a := p1
	points := [][]int{p1, p2, p3, p4}
	length := 0
	var b []int
	for i := 1; i < 4; i++ {
		if a[0] == points[i][0] {
			length = abs(points[i][1] - p1[1])
			b = points[i]
		}
	}
	if length == 0 {
		return false
	}
	diff := 0
	for i := 1; i < 4; i++ {
		if a[1] == points[i][1] && a[0] != points[i][0] {
			diff = p1[1] - points[i][0]
		}
	}
	if diff == 0 {
		return false
	}
	for i := 1; i < 4; i++ {
		if b[1] == points[i][1] && b[0] != points[i][0] {
			return diff == b[0]-points[i][0]
		}
	}
	return false
}

func main() {
	validSquare([]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1})
}
