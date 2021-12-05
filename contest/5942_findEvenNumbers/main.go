package main

import "sort"

type void struct {
}

var member void

func findEvenNumbers(digits []int) []int {
	ans := sort.IntSlice{}
	var dict map[int]void
	dict = make(map[int]void)
	size := len(digits)
	for i := 0; i < size; i++ {
		// 最后一位
		if digits[i]%2 != 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if j == i {
				continue
			}
			for k := 0; k < size; k++ {
				if digits[k] == 0 {
					continue
				}
				if k == i || k == j {
					continue
				}
				num := digits[k]*100 + digits[j]*10 + digits[i]
				if _, ok := dict[num]; ok {
					continue
				}
				dict[num] = member
				ans = append(ans, num)
			}
		}
	}
	sort.Sort(ans)
	return ans
}

func main() {
	findEvenNumbers([]int{2, 1, 3, 0})
}
