package main

import (
	"fmt"
)

func main() {
	const size = 10

	// 1st experiment array
	var arr [size]int
	arr = [size]int{10, 20, 30, 40, 50}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Array position ", i, " have data ", arr[i], " at address ", &arr[i])
	}
	fmt.Println("length of the array ", len(arr))
	fmt.Println("capacity of the array ", cap(arr))

	// 2nd experiment array
	var arr1 *[size]int
	arr1 = &[size]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr1); i++ {
		fmt.Println("Array position ", i, " have data ", arr1[i], " at address ", &arr1[i])
	}
	fmt.Println("length of the array ", len(arr1))
	fmt.Println("capacity of the array ", cap(arr1))

	// 3rd experiment array
	var arr2 = new([size]int)
	arr2 = &[size]int{23, 34, 45, 56, 65, 21}
	for i := 0; i < len(arr1); i++ {
		fmt.Println("Slice position ", i, " have data ", arr2[i], " at address ", &arr2[i])
	}
	fmt.Println("length of the array ", len(arr2))
	fmt.Println("capacity of the array ", cap(arr2))

	// 4th experiment slice
	// Declare Slice using Make
	var slice = make([]int, 5, size)
	// slice = []int{7, 64, 34, 234, 456}
	slice[1] = 10
	slice[4] = 20
	slice = slice[:cap(slice)]
	slice[5] = 20
	for i := 0; i < len(slice); i++ {
		fmt.Println("Slice position ", i, " have data ", slice[i], " at address ", &slice[i])
	}
	fmt.Println("length of the array ", len(slice))
	fmt.Println("capacity of the array ", cap(slice))

	// 5th experiment slice
	// Create Empty Slice
	var slice1 []int
	fmt.Printf("value %v and Type %T", slice1, slice1)

	// 6th experiment slice
	// Initialize Slice with values using a Slice Literal
	var slice2 = []int{34, 54, 657, 234, 456}
	var slice3 = []string{"ashish", "abhay", "pooja", "trishiak", "neha"}
	for i := 0; i < len(slice2); i++ {
		fmt.Println("Slice position ", i, " have data ", slice2[i], " at address ", &slice2[i])
	}
	fmt.Println("length of the array ", len(slice2))
	fmt.Println("capacity of the array ", cap(slice2))

	for i := 0; i < len(slice3); i++ {
		fmt.Println("Slice position ", i, " have data ", slice3[i], " at address ", &slice3[i])
	}
	fmt.Println("length of the array ", len(slice2))
	fmt.Println("capacity of the array ", cap(slice2))

	// 7th experiment slice
	// Declare Slice using new Keyword
	var slice4 = new([size]int)[:size]
	for i := 0; i < len(slice4); i++ {
		fmt.Println("Slice position ", i, " have data ", slice4[i], " at address ", &slice4[i])
	}
	fmt.Println("length of the array ", len(slice4))
	fmt.Println("capacity of the array ", cap(slice4))
}
