package iteration

import "strings"

//Addtional variants of the loop are describe here: https://gobyexample.com/for

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
