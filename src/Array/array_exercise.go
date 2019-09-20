package main

import (
	"fmt"
)

func main() {
	a := [...]string{"USA", "China", "India", "France"}
	var b = a
	b[0] = "SingSing"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
	num := [...]int{1, 2, 3, 5, 8}
	fmt.Println("before change", num)
	//changeLocal(num)
	fmt.Println("After changeLocal", num)
	ergodicArray()
	rangeArray()
}

func changeLocal(num [5]int) {
	num[0] = 50
	fmt.Println("inside function", num)
}

func ergodicArray() {
	a := [...]float64{37.3, 22, 33, 22.09, 5}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}

func rangeArray() {
	a := [...]float64{37.3, 22, 33, 22.09, 5}
	sum := float64(0)
	for _, v := range a {
		sum += v
		fmt.Println("The new element is ", sum)
	}
}
