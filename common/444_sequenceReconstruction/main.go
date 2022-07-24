package main

type void struct{}

var member void

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	// 入度出度
	// 使用线段表
	out := map[int][]int{}
	in := map[int]int{}
	set := map[int]void{}
	for i := 0; i < len(sequences); i++ {
		set[sequences[i][0]] = member
		for j := 0; j < len(sequences[i])-1; j++ {
			outPoint := j + 1
			inPoint := j
			if _, ok := out[sequences[i][outPoint]]; !ok {
				out[sequences[i][outPoint]] = []int{}
			}
			out[sequences[i][inPoint]] = append(out[sequences[i][inPoint]], sequences[i][outPoint])
			if _, ok := in[sequences[i][outPoint]]; !ok {
				in[sequences[i][outPoint]] = 0
			}
			in[sequences[i][outPoint]]++
			set[sequences[i][outPoint]] = member
		}
	}
	if len(nums) > len(set) {
		return false
	}
	povit := -1
	for key, _ := range set {
		if _, ok := in[key]; !ok {
			if povit != -1 {
				return false
			}
			povit = key
		}
	}

	// 若有分叉则直接否
	// 接着判断长度
	for len(in) > 0 {
		count := false
		pres := povit
		for i := 0; i < len(out[pres]); i++ {
			in[out[pres][i]]--
			if in[out[pres][i]] == 0 {
				delete(in, out[pres][i])
				povit = out[pres][i]
				if count {
					return false
				}
				count = true
			}
		}
	}
	return true
}
