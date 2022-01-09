package main

type void struct{}

var member void

var bits []int
var bitsMap map[int]bool

func wordCount(startWords []string, targetWords []string) int {
	// 使用位运算
	if len(bits) == 0 {
		pres := 1
		bitsMap = make(map[int]bool)
		for i := 0; i < 26; i++ {
			bits = append(bits, pres)
			bitsMap[pres] = true
			pres *= 2
		}
	}
	startSize := len(startWords)
	targetSize := len(targetWords)
	startBits := make(map[int]void)
	piviot := "a"[0]
	for i := 0; i < startSize; i++ {
		bit := 0
		wordSize := len(startWords[i])
		for j := 0; j < wordSize; j++ {
			bit |= bits[startWords[i][j]-piviot]
		}
		startBits[bit] = member
	}
	targetBits := []int{}
	for i := 0; i < targetSize; i++ {
		bit := 0
		wordSize := len(targetWords[i])
		for j := 0; j < wordSize; j++ {
			bit |= bits[targetWords[i][j]-piviot]
		}
		targetBits = append(targetBits, bit)
	}
	// 开始比较
	ans := 0
	for i := 0; i < targetSize; i++ {
		target := targetBits[i]
		for j := 0; j < 26; j++ {
			if bits[j]|target == target {
				if _, ok := startBits[target^bits[j]]; ok {
					ans += 1
					break
				}
			}
		}
	}
	return ans
}

func main() {
	// wordCount([]string{"uh"}, []string{"u", "hur", "k", "b", "u", "yse", "giqoy", "lni", "olqb", "nemc"})
	wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"})
	// wordCount([]string{"ab", "a"}, []string{"abc", "abcd"})
}
