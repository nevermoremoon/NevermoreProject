package main

import "fmt"

func subtacOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= -2
	}
}

func main() {
	nos := []int{5, 6, 9}
	fmt.Println("slice before function call", nos)
	subtacOne(nos)
	fmt.Println("slice after function call", nos)
}

//切片包含长度、容量和指向数组第零个元素的指针。当切片传递给函数时，即使它通过值传递，
// 指针变量也将引用相同的底层数组。
// 因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见
