package main

type void struct{}

var member void

func distributeCandies(candyType []int) int {
	// 使用哈希表空键作为值
	size := len(candyType)
	half := size >> 1
	set := make(map[int]void)
	for i := 0; i < size; i++ {
		set[candyType[i]] = member
	}
	if len(set) > half {
		return half
	}
	return len(set)
}
