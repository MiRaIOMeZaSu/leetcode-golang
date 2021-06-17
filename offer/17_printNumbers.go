package main

func printNumbers(n int) []int {
	pivot := 0
	temp := 1
	for i := 0; i < n; i++ {
		pivot += 9 * temp
		temp *= 10
	}
	var ret []int
	for i := 1; i <= pivot; i++ {
		ret = append(ret, i)
	}
	return ret
}
