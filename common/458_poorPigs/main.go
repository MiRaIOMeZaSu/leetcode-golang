package main

const (
	INT_MAX = int(^uint32(0) >> 1)
)

var bits []int

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	// 使用位运算
	// 4buckets2pigs一次测完的情况 - 小猪a:1,2 小猪b: 2,3
	// 1000 buckets 5 pigs 4次相当于 251 buckets 5 pigs 1次(256 = 2^8)
	// 32 16 8 4(1000/32=31.5 32/16=2(2:3))
	// 1000-5: 167-4 42-3 14-2
	// 2^5: 32 2^20: 1048576 2^9: 512
	if len(bits) == 0 {
		bits = append(bits, 1)
		for i := 1; i <= 10+1; i++ {
			bits = append(bits, bits[i-1]<<1)
		}
	}
	if buckets < 2 {
		return 0
	}
	times := minutesToTest / minutesToDie
	if times+1 >= buckets {
		return 1
	}
	left := 2
	right := 10
	ans := right
	for left <= right {
		// 判断能否满足要求
		mid := (left + right) >> 1
		if statify(buckets, times, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func statify(buckets int, times int, pigs int) bool {
	// 为什么会满足条件
	currBuckets := buckets
	lastBuckets := buckets
	currTimes := times
	currPigs := pigs
	// 测试次数
	// times超大情况下如何?
	for currTimes-1 > currPigs {
		buckets -= bits[currPigs-1]
		currTimes--
	}
	for i := 0; i < currTimes-1 && currPigs > 0; i++ {
		if bits[currPigs] >= currBuckets {
			return true
		} else {
			lastBuckets = currBuckets
			currBuckets /= bits[currPigs]
			if currBuckets*bits[currPigs] < lastBuckets {
				currBuckets += 1
			}
			currPigs--
		}
	}
	if times == 1 {
		return bits[pigs] >= buckets
	}
	return false
}

func main() {
	if len(bits) == 0 {
		bits = append(bits, 1)
		for i := 1; i <= 10+1; i++ {
			bits = append(bits, bits[i-1]<<1)
		}
	}
	// fmt.Print(poorPigs(301,
	// 	12,
	// 	95))
	statify(301, 7, 3)
}
