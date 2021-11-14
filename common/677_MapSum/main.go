package main

import "fmt"

type Tree struct {
	sum   int
	nodes [26]*Tree
}

type MapSum struct {
	// 使用字典树
	stringMap map[string]int
	root      *Tree
}

func Constructor() MapSum {
	var mapSum MapSum
	mapSum.root = &Tree{0, [26]*Tree{}}
	mapSum.stringMap = make(map[string]int)
	return mapSum
}

func (this *MapSum) Insert(key string, val int) {
	curr := this.root
	divide := val
	if value, ok := this.stringMap[key]; ok {
		divide -= value
	}
	this.stringMap[key] = val
	for _, ch := range key {
		index := ch - 'a'
		if curr.nodes[index] == nil {
			curr.nodes[index] = &Tree{0, [26]*Tree{}}
		}
		curr = curr.nodes[index]
		curr.sum += divide
	}
}

func (this *MapSum) Sum(prefix string) int {
	curr := this.root
	for _, ch := range prefix {
		index := ch - 'a'
		if curr.nodes[index] == nil {
			return 0
		}
		curr = curr.nodes[index]
	}
	return curr.sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

func main() {
	obj := Constructor()
	obj.Insert("apple", 3)
	fmt.Print(obj.Sum("ap"))
	obj.Insert("app", 2)
	fmt.Print(obj.Sum("ap"))
}
