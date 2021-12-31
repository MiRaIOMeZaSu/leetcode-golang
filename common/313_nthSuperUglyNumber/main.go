package main

import "sort"

type void struct{}

var member void

var size int
var arr sort.IntSlice
var set map[int]void

func solve(primes *[]int, index int, curr int) {
	if _, ok := set[curr]; index >= size || curr > (*primes)[size-1] || ok {
		return
	}
	arr = append(arr, curr)
	set[curr] = member
	solve(primes, index+1, curr*(*primes)[index])
	solve(primes, index+1, curr)
}

func nthSuperUglyNumber(n int, primes []int) int {
	// 假如使用优先队列法,应该使得新入队的是当前最小的
	// 计算素数的数量?变化为前缀法?
	size = len(primes)
	set = make(map[int]void)
	ans := 1
	arr = sort.IntSlice{}
	solve(&primes, 0, 1)
	sort.Sort(arr)
	
	return ans
}
