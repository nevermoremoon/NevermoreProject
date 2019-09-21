package main

import "fmt"

//切片是动态的，使用 append 可以将新元素追加到切片上。
// append 函数的定义是 func append（s[]T，x ... T）[]T
func main() {
	cars := []string{"Ferrai", "Honda", "Ford"}
	fmt.Println("cars", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars", cars, "has nes length", len(cars), "and capacity", cap(cars))
	appendNil()
	appendMoreSlice()
}

//当新的元素被添加到切片时，会创建一个新的数组。
// 现有数组的元素被复制到这个新数组中，并返回这个新数组的新切片引用。
// 现在新切片的容量是旧切片的两倍

//append nil
func appendNil() {
	var names []string
	if names == nil {
		fmt.Println("Start to Append")
		names = append(names, "lala", "hihi", "hehe")
		fmt.Println(names, len(names), cap(names))
	}
}

func appendMoreSlice() {
	veggtables := []string{"potatas", "tomato", "brjinal"}
	fruits := []string{"oranges", "watermelon"}
	var food = append(veggtables, fruits...)
	fmt.Println(food)
}
