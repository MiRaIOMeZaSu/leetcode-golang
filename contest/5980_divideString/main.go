package main

func divideString(s string, k int, fill byte) []string {
	ans := []string{}
	size := len(s)
	needle := k - size%k
	temp := ""
	for i := 0; i < size; i++ {
		temp += string(s[i])
		if len(temp) == k {
			ans = append(ans, temp)
			temp = ""
		}
	}
	for i := 0; needle != k && i < needle; i++ {
		temp += string(fill)
	}
	if len(temp) == k {
		ans = append(ans, temp)
	}
	return ans
}
func main() {
	divideString("abcdefghij",
		3,
		'x')
}
