package main

func toLowerCase(s string) string {
	var ans []byte
	for i := 0; i < len(s); i++ {
		index := s[i] - 'A'
		if index >= 0 && index < 26 {
			ans = append(ans, 'a'+index)
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
