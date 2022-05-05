package main

import "fmt"

func main() {
	// when the slice capcity will full
	// it double the capcity every time
	var slice []int
	var count int = 100
	for i := 0; i < count; i++ {
		slice = append(slice, i+count)
		fmt.Println("length of the array ", len(slice))
		fmt.Println("capacity of the array ", cap(slice))
		fmt.Println()
	}
	fmt.Println(slice)

}
