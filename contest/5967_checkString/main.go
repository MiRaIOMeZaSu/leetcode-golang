package main

func checkString(s string) bool {
	i := 0
	a := "a"[0]
	b := "b"[0]
	size := len(s)
	for i < size && s[i] == a {
		i++
	}
	for i < size && s[i] == b {
		i++
	}
	return i == size
}

func main() {
	checkString("aaabbb")
}
