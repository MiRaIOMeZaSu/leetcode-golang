package main

func countPoints(rings string) int {
	var table [10][3]int
	var dict map[byte]int
	dict = map[byte]int{}
	dict['R'] = 0
	dict['G'] = 1
	dict['B'] = 2
	for i := 0; i < len(rings); i += 2 {
		table[rings[i+1]-'0'][dict[rings[i]]] = 1
	}
	ans := 0
	flag := true
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			if table[i][j] == 0 {
				flag = false
				break
			}
		}
		if flag {
			ans++
		} else {
			flag = true
		}
	}
	return ans
}
