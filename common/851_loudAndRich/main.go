package main

import (
	"container/list"
	"fmt"
)

type void struct{}

var member void

type Queue struct {
	inner *list.List
}

func Constructor() Queue {
	q := Queue{list.New()}
	return q
}

func (this Queue) add(val int) {
	this.inner.PushBack(val)
}
func (this Queue) poll() list.Element {
	front := this.inner.Front()
	this.inner.Remove(front)
	return *front
}

func (this Queue) len() int {
	return this.inner.Len()
}

func loudAndRich(richer [][]int, quiet []int) []int {
	// 依旧是简单的拓扑排序
	n := len(quiet)
	graph := [][]int{}
	inCount := []int{}
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, i)
		inCount = append(inCount, 0)
		graph = append(graph, []int{})
	}
	for i := 0; i < len(richer); i++ {
		// b比a更有钱,从有钱到没钱的(不少于)
		a := richer[i][0]
		b := richer[i][1]
		inCount[b]++
		graph[a] = append(graph[a], b)
	}
	q := Constructor()
	for i := 0; i < n; i++ {
		if inCount[i] == 0 {
			q.add(i)
		}
	}
	visited := make(map[int]void)
	for q.len() != 0 {
		pivot := q.poll().Value.(int)
		visited[pivot] = member
		for i := 0; i < len(graph[pivot]); i++ {
			next := graph[pivot][i]
			if _, ok := visited[next]; !ok {
				inCount[next]--
				if quiet[ans[pivot]] < quiet[ans[next]] {
					ans[next] = ans[pivot]
				}
				if inCount[next] == 0 {
					q.add(next)
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Print(loudAndRich([][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}},
		[]int{3, 2, 5, 4, 6, 1, 7, 0}))
}
