package main

import "fmt"

func modify(arr *[3]int) {
	(*arr)[0] = 10
}
func main() {
	a := [3]int{1, 3, 4}
	modify(&a)
	fmt.Println(a)
}
