package main

func minMoves(target int, maxDoubles int) int {
	// 如何保证行动次数最小
	ans := 0
	curr := target
	for maxDoubles > 0 && curr != 1 {
		if curr%2 != 0 {
			ans++
			curr--
		} else {
			maxDoubles--
			ans++
			curr /= 2
		}
	}
	ans += curr - 1
	return ans
}

func main() {
	// minMoves(5, 0)
	// minMoves(19, 2)
	println(minMoves(10, 4))
}
