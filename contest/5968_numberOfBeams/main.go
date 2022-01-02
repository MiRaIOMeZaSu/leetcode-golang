package main

func numberOfBeams(bank []string) int {
	// 逐层下去
	arr := []int{}
	ans := 0
	size := len(bank)
	for i := 0; i < size; i++ {
		temp := 0
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				temp++
			}
		}
		if temp > 0 {
			arr = append(arr, temp)
		}
	}
	// 开始逐层往下
	for i := 0; i < len(arr)-1; i++ {
		ans += arr[i] * arr[i+1]
	}
	return ans
}
