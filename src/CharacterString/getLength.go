package main

import (
	"fmt"
	"unicode/utf8"
)

func getStringLength(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutates(ss []rune) string {
	ss[0] = 'a'
	return string(ss)
}

func main() {
	word1 := "Señor"
	getStringLength(word1)
	word2 := "Pets"
	getStringLength(word2)
	fmt.Println(mutates([]rune(word1)))
}
