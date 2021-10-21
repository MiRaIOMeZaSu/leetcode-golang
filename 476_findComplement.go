package main

func findComplement(num int) int {
	i := 31
	pivit := int(^uint32(0)>>2) + 1
	for ; i >= 0; i-- {
		if (num | pivit) == num {
			break
		}
		pivit >>= 1
	}
	pivit = (pivit << 1) - 1
	return pivit ^ num
}

func main() {
}
