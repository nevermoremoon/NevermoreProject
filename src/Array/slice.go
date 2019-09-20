package main

import "fmt"

func main() {
	a := [...]int{1, 3, 44, 55, 666, 6662}
	var b []int = a[1:4]
	fmt.Println(b)
	c := []int{6, 7, 8}
	fmt.Println(c)
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	fmt.Println(nums1)
	fmt.Println(nums2)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
}
