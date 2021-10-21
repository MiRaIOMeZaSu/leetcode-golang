package main

import "fmt"

func findComplement(num int) int {
	return -1 ^ num
}

func main() {
	fmt.Print(findComplement(5))
}
