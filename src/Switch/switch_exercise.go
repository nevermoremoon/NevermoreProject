package main

import "fmt"

func main() {
	//finger := 4
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Anthoer Ring")
	default:
		fmt.Println("Please input right number of finger!")
	}
	moreCase()
	simpleSwitch()
	fallThroughExcise()
}

func moreCase() {
	letter := "i"
	switch letter {
	case "a", "o", "i", "e":
		fmt.Println("Vetory!")
	default:
		fmt.Println("The chactaor is not in letter!")
	}
}
func simpleSwitch() {
	num := 75
	switch {

	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50！")
	case num > 50 && num <= 100:
		fmt.Println("num is greater than 50 and less than 100!")
	case num > 100:
		fmt.Println("num is  greater than 100!")

	}
}

func getNumber() int {
	num := 4 * 20
	return num
}

func fallThroughExcise() {
	switch num := getNumber(); {
	case num < 50:
		fmt.Printf("The %d is less than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("The %d is less than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("The %d is less than 200\n", num)

	}
}

//fallthrough 不会去判断下一个case的值是否为true
