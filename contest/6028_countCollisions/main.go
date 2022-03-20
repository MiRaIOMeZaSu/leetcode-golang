package main

func countCollisions(directions string) int {
	left := 0
	lPivot := "L"[0]
	rPivot := "R"[0]
	sPivot := "S"[0]
	right := len(directions) - 1
	for ; left <= right && directions[left] == lPivot; left++ {
	}
	for ; right >= 0 && directions[right] == rPivot; right-- {
	}
	ans := 0
	for i := left; i <= right; i++ {
		if directions[i] != sPivot {
			ans++
		}
	}
	return ans
}
