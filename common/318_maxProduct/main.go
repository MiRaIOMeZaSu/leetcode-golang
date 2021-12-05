package main

import (
	"fmt"
	"sort"
)

var bits [26]int

type MyWord struct {
	length int
	val    int
}

type MyWords []*MyWord

func (words MyWords) Len() int {
	return len(words)
}

func (words MyWords) Less(i, j int) bool {
	if words[i].length != words[j].length {
		return words[i].length > words[j].length
	}
	return words[i].val > words[j].val
}

func (words MyWords) Swap(i, j int) {
	words[i], words[j] = words[j], words[i]
}

func getVal(word string) int {
	ans := 0
	for i := 0; i < len(word); i++ {
		ans |= bits[word[i]-'a']
	}
	return ans
}

func maxProduct(words []string) int {
	bits[0] = 1
	for i := 1; i < 26; i++ {
		bits[i] = bits[i-1] << 1
	}
	// 遍历
	size := len(words)
	var myWords MyWords
	for i := 0; i < size; i++ {
		myWords = append(myWords, &MyWord{len(words[i]), getVal(words[i])})
	}
	sort.Sort(myWords)
	ans := 0
	for i := 0; i < size; i++ {
		if i+1 < size && myWords[i].length*myWords[i+1].length < ans {
			break
		}
		for j := i + 1; j < size; j++ {
			temp := myWords[i].length * myWords[j].length
			if temp <= ans {
				break
			}
			if myWords[i].val|myWords[j].val == myWords[i].val+myWords[j].val {
				// 更新
				ans = temp
			}
		}
	}
	return ans
}

func main() {
	fmt.Print(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}
