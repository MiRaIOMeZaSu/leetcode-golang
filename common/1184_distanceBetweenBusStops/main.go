package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	total := 0
	for i := 0; i < len(distance); i++ {
		total += distance[i]
	}
	_start := min(start, destination)
	_destination := max(start, destination)
	length := 0
	for i := _start; i < _destination; i++ {
		length += distance[i]
	}
	return min(length, total-length)
}
