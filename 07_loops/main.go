package main

import "fmt"

func main() {
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for j := 1; j <= 10; j++ {
	// 	fmt.Printf("Number %d\n", j)
	// }

	// Fizzbuzz

	for j := 1; j <= 100; j++ {
		if j % 15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if j % 3 == 0 {
			fmt.Println("Fizz")
		} else if j % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}
}