package main

func findJudge(n int, trust [][]int) int {
	in := []int{}
	out := []int{}
	for i := 0; i < n+1; i++ {
		in = append(in, 0)
		out = append(out, 0)
	}
	size := len(trust)
	for i := 0; i < size; i++ {
		in[trust[i][1]]++
		out[trust[i][0]]++
	}
	for i := 1; i <= n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i
		}
	}
	return -1
}

func main() {
	findJudge(2, [][]int{{1, 2}})
}
