package main

import "fmt"

func main() {
	x := 5
	y := 5

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if  x > y{
		fmt.Printf("%d is less than %d\n", y, x)
	} else {
		fmt.Printf("%d is equal to %d\n", x, y)
	}

	color := "blue"

	if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue")
	}

	// Switch

	switch color {
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue")
	}
}