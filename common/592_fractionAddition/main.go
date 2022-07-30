package main

type void struct {
}

var member void

func fractionAddition(expression string) string {
	chars := []rune(expression)
	set := map[rune]void{[]rune("+")[0]: member, []rune("-")[0]: member}
	for i := 0; i < len(chars); {
		if _, ok := set[chars[i]]; ok {
			// 此时可以分割
			i++
		}
		_, ok := set[chars[i]]
		for !ok {

			i++
			_, ok = set[chars[i]]
		}
	}
}
