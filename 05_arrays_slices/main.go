package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2] string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2] string {"Apple", "Orange"}

	// Slice
	// Slice is a pointer pointing to the 1st element, so if a variable copies it and changed, the original slice is also changed. Array doesn't change.
	fruitArr := [] string {"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[1:3])

	fruitArr = append(fruitArr, "Banana")
}