package main

type void struct{}

var member void

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 拓扑排序
	// 邻接表
	// 只要不存在环必然满足
	var graph [][]int
	inCount := []int{}
	for i := 0; i < numCourses; i++ {
		graph = append(graph, []int{})
		inCount = append(inCount, 0)
	}
	for i := 0; i < len(prerequisites); i++ {
		a := prerequisites[i][0]
		b := prerequisites[i][1]
		graph[a] = append(graph[a], b)
		inCount[b]++
	}
	var visited map[int]void
	visited = make(map[int]void)
	starter := []int{}
	for i := 0; i < numCourses; i++ {
		if inCount[i] == 0 {
			starter = append(starter, i)
		}
	}
	if len(starter) == 0 {
		return false
	}
	// 实现队列: slice/list
	// 计算入度
	queue := []int{}
	queue = append(queue, starter...)
	for len(queue) != 0 {
		// 出队
		pivot := queue[0]
		queue = queue[1:]
		visited[pivot] = member
		for j := 0; j < len(graph[pivot]); j++ {
			next := graph[pivot][j]
			if _, ok := visited[next]; !ok {
				// 未访问过
				inCount[next]--
				if inCount[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}
	for i := 0; i < numCourses; i++ {
		if inCount[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	canFinish(2, [][]int{{0, 1}})
}
