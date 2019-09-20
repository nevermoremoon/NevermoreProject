package main

import "fmt"

func printarray(a [3][2]string) {
	for _, v := range a {
		fmt.Printf("%s", v)
	}
	fmt.Printf("\n")
}
func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}
