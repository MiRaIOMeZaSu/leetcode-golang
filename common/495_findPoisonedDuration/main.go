package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLeft(arr *[]int, pivot int, leftIndex int) int {
	// 在数组中寻找值小于或等于pivot的值
	rightIndex := len(*arr) - 1
	ans := leftIndex
	for leftIndex <= rightIndex {
		mid := (leftIndex + rightIndex) >> 1
		if (*arr)[mid] < pivot {
			ans = max(mid, ans)
			leftIndex = mid + 1
		} else if (*arr)[mid] > pivot {
			rightIndex = mid - 1
		} else {
			return mid - 1
		}
	}
	return ans
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	// 对于排序序列,通过二分法获得指定值靠左的值
	// 建立数字到索引的映射?
	ans := 0
	for i := 0; i < len(timeSeries); i++ {
		startMoment := timeSeries[i]
		endMoment := timeSeries[i] + duration - 1
		leftStart := findLeft(&timeSeries, endMoment+1, i)
		for leftStart > i {
			// 可以连续
			i = leftStart
			endMoment = timeSeries[i] + duration - 1
			leftStart = findLeft(&timeSeries, endMoment+1, i)
		}
		ans += endMoment - startMoment + 1
	}
	return ans
}

func main() {
	findPoisonedDuration([]int{1, 2}, 2)
}
