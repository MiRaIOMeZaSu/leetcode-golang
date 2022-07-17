package main

type void struct {
}

var member void

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func arrayNesting(nums []int) int {
	// 动态规划
	// 记录数字的下标
	var dict map[int]int = make(map[int]int)
	var visited map[int]void = make(map[int]void)
	lenght := len(nums)
	for i := 0; i < lenght; i++ {
		dict[nums[i]] = i
	}
	ans := 1
	for i := 0; i < lenght; i++ {
		curr := nums[i]
		if _, ok := visited[curr]; ok {
			continue
		}
		visited[curr] = member
		list := []int{curr}
		var currVisited map[int]void = make(map[int]void)
		currVisited[curr] = member
		next := nums[curr]
		_, ok := currVisited[next]
		for !ok {
			currVisited[next] = member
			visited[next] = member
			list = append(list, next)
			next = nums[next]
			_, ok = currVisited[next]
		}
		ans = max(ans, len(list))
	}
	return ans
}

func main() {
	list := []int{5, 4, 0, 3, 1, 6, 2}
	print(arrayNesting(list))
}
