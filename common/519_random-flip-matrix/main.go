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
	curr    int
	m       int
	n       int
	border  int
	rest    int
}

func Constructor(m int, n int) Solution {
	rand.Seed(time.Now().Unix())
	return Solution{make(map[int]void), 0, m, n, m * n, m * n}
}

func (this *Solution) Flip() []int {
	currPercent := 1.0 / float64(this.rest)
	for {
		for i := this.curr; i < this.border; i++ {
			if _, ok := this.visited[i]; ok {
				continue
			}
			if rand.Float64() < currPercent {
				this.curr = (i + 1) % this.border
				this.rest--
				this.visited[i] = member
				return indexToPosition(this.n, i)
			}
			// 继续循环,更新currPercent
			currPercent = 1.0 / float64(this.rest)
		}
		this.curr = 0
	}
}

func (this *Solution) Reset() {
	this.rest = this.border
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
