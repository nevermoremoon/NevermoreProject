package main

import "fmt"

type employ struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp8 := &employ{"Sam", "Anderson", 19, 30000}
	fmt.Println("Firstname is :", (*emp8).firstName)
	fmt.Println("Age is :", (*emp8).age)
}
