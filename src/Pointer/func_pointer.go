package main

import "fmt"

func main() {
	a := 58
	fmt.Println("value before function call is:", a)
	var b *int = &a
	change(b)
	fmt.Println("value after function call is:", a)
}

func change(val *int) {
	*val = 55
}
