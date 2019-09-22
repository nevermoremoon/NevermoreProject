package main

import "fmt"

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printCharacter(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) //在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间,
		//如果字符串中有单个字符超过1个字节，就会打印报错
	}
}

func printChars(s string) {
	runes := []rune(s) //代码点无论占用多少个字节，都可以用一个 rune 来表示
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printByesAndChars(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts byte %d\n", rune, index)
	}
}

func main() {
	printBytes("Hello World")
	fmt.Printf("\n")
	printCharacter("Hello World")
	names := "Señor"
	printBytes(names)
	fmt.Printf("\n")
	printCharacter(names)
	fmt.Printf("\n")
	printChars(names)
	fmt.Printf("\n")
	printByesAndChars(names)
}
