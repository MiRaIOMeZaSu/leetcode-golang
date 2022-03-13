package main

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	ans := 0
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}
	for i := 0; i < len(dig); i++ {
		graph[dig[i][0]][dig[i][1]] = 1
	}
	for i := 0; i < len(artifacts); i++ {
		// 工件为矩形
		flag := true
		for x := artifacts[i][0]; flag && x <= artifacts[i][2]; x++ {
			for j := artifacts[i][1]; flag && j <= artifacts[i][3]; j++ {
				if graph[x][j] == 0 {
					flag = false
					break
				}
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}

func main() {
	digArtifacts(6, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}})
	digArtifacts(6,
		[][]int{{0, 2, 0, 5}, {0, 1, 1, 1}, {3, 0, 3, 3}, {4, 4, 4, 4}, {2, 1, 2, 4}},
		[][]int{{0, 2}, {0, 3}, {0, 4}, {2, 0}, {2, 1}, {2, 2}, {2, 5}, {3, 0}, {3, 1}, {3, 3}, {3, 4}, {4, 0}, {4, 3}, {4, 5}, {5, 0}, {5, 1}, {5, 2}, {5, 4}, {5, 5}})
}
