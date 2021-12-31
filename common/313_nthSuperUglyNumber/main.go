package main

import (
	"container/heap"
)

type Element struct {
	Val     int
	Index   int
	Pointer int
}

// 优先队列，小顶堆
type PriorQ []Element

// 以下方法为heap.Interface的模板方法，需要熟练掌握
func (p PriorQ) Len() int {
	return len(p)
}

func (p PriorQ) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p PriorQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorQ) Push(x interface{}) {
	*p = append(*p, x.(Element))
}

func (p *PriorQ) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

func nthSuperUglyNumber(n int, primes []int) int {
	// 假如使用优先队列法,应该使得新入队的是当前最小的
	// 计算素数的数量?变化为前缀法?
	p := &PriorQ{}
	heap.Init(p)
	size := len(primes)
	for i := 0; i < size; i++ {
		heap.Push(p, Element{primes[i], i, 0})
	}
	ugly := []int{1}
	for len(ugly) < n {
		head := heap.Pop(p)
		element := head.(Element)
		val := element.Val
		if val != ugly[len(ugly)-1] {
			ugly = append(ugly, val)
		}
		element.Pointer = element.Pointer + 1
		element.Val = ugly[element.Pointer] * primes[element.Index]
		heap.Push(p, element)
	}
	return ugly[n-1]
}

func main() {
	nthSuperUglyNumber(12, []int{2, 7, 13, 19})
	nthSuperUglyNumber(1, []int{2, 3, 5})
}
