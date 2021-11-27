package main

import (
	"fmt"
	"math/rand"
	"time"
)

func indexToPosition(n, index int) []int {
	return []int{index / n, index % n}
}

type void struct{}

var member void

type Solution struct {
	// 使用集合存储已经反转过的数字
	visited map[int]void
	n       int
	border  int
}

func Constructor(m int, n int) Solution {
	rand.Seed(time.Now().Unix())
	return Solution{make(map[int]void), n, m * n}
}

func (this *Solution) Flip() []int {
	for {
		i := rand.Intn(this.border)
		if _, ok := this.visited[i]; ok {
			continue
		}
		this.visited[i] = member
		return indexToPosition(this.n, i)
	}
}

func (this *Solution) Reset() {
	this.visited = make(map[int]void)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */

func main() {
	obj := Constructor(3, 1)
	fmt.Print(obj.Flip())
	fmt.Print(obj.Flip())
	fmt.Print(obj.Flip())
	obj.Reset()
	fmt.Print(obj.Flip())
}
