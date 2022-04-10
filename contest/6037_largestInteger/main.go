package main

import "sort"

func largestInteger(num int) int {
	isOdd := []bool{}
	listOdd := sort.IntSlice{}
	listEven := sort.IntSlice{}
	for i := 0; num > 0; i++ {
		curr := num % 10
		num /= 10
		if curr%2 == 0 {
			isOdd = append(isOdd, false)
			listEven = append(listEven, curr)
		} else {
			isOdd = append(isOdd, true)
			listOdd = append(listOdd, curr)
		}
	}
	sort.Sort(listOdd)
	sort.Sort(listEven)
	oddIndex := 0
	evenIndex := 0
	index := 0
	ans := 0
	for i := 1; index < len(isOdd); i *= 10 {
		if isOdd[index] {
			ans += listOdd[oddIndex] * i
			oddIndex++
		} else {
			ans += listEven[evenIndex] * i
			evenIndex++
		}
		index++
	}
	return ans
}

func main() {
	largestInteger(1234)
}
