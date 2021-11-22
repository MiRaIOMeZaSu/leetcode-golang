package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	size  int
	proto []int
	curr  []int
}

func swap(i, j int, nums []int) {
	nums[j], nums[i] = nums[i], nums[j]
}

func Constructor(nums []int) Solution {
	solution := Solution{}
	solution.proto = nums
	solution.curr = append(solution.curr, nums...)
	solution.size = len(nums)
	return solution
}

func (this *Solution) Reset() []int {
	return this.proto
}

func (this *Solution) Shuffle() []int {
	// 开始插入
	// 200个数字
	// 5*10^4次调用
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < this.size; i++ {
		swap(i, r.Intn(this.size), this.curr)
	}
	return this.curr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
