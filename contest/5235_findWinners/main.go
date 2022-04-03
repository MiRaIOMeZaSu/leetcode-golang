package main

import "sort"

type void struct{}

var member void

func findWinners(matches [][]int) [][]int {
	winners := sort.IntSlice{}
	littleWinners := sort.IntSlice{}
	var loseTime map[int]int
	loseTime = make(map[int]int)
	for i := 0; i < len(matches); i++ {
		loser := matches[i][1]
		for j := 0; j < 2; j++ {
			if _, ok := loseTime[matches[i][j]]; !ok {
				loseTime[matches[i][j]] = 0
			}
		}
		loseTime[loser] += 1
	}
	for key, val := range loseTime {
		if val == 0 {
			winners = append(winners, key)
		} else if val == 1 {
			littleWinners = append(littleWinners, key)
		}
	}
	winners.Sort()
	littleWinners.Sort()
	return [][]int{winners, littleWinners}
}
