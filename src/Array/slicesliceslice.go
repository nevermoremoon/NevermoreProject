package main

import "fmt"

func main() {
	lug := [][]string{
		{"c", "c++"},
		{"js"},
		{"go", "ruby"},
	}
	for _, v1 := range lug {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
