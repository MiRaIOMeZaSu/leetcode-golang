package main

const (
	INT_MAX = int(^uint((0)) >> 1)
)

func giveGem(gem []int, operations [][]int) int {
	for i := 0; i < len(operations); i++ {
		temp := gem[operations[i][0]] >> 1
		gem[operations[i][0]] -= temp
		gem[operations[i][1]] += temp
	}
	min := INT_MAX
	max := 0
	for i := 0; i < len(gem); i++ {
		if min > gem[i] {
			min = gem[i]
		}
		if max < gem[i] {
			max = gem[i]
		}
	}
	return max - min
}
