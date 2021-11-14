package main

type Tree struct {
	val   int
	sum   int
	nodes [26]*Tree
}

type MapSum struct {
	// 使用字典树
	root *Tree
}

func Constructor() MapSum {
	var mapSum MapSum
	mapSum.root = &Tree{0, 0, [26]*Tree{}}
	return mapSum
}

func (this *MapSum) Insert(key string, val int) {
	curr := this.root
	for _, ch := range key {
		index := ch - 'a'
		if curr.nodes[index] == nil {
			curr.nodes[index] = &Tree{0, 0, [26]*Tree{}}
		}
		curr.sum += val
		curr = curr.nodes[index]
	}
	curr.val = val
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
