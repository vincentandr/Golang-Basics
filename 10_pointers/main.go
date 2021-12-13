package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, *b)
	fmt.Println(a, b)
	fmt.Println(*&a)
	fmt.Printf("%T\n", b)

	*b = 10
	fmt.Println(a, *b)
}