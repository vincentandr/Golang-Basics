package main

import "fmt"

func greeting(name string) (string, error) {
	return "Hello " + name, nil
}

func getSum(num1 int, num2 int) (int, error) {
	return num1 + num2, nil
}

func main() {
	fmt.Println(greeting("Vincent"))
	fmt.Println(getSum(1, 4))
	
	// Anonymous function, only called exactly 1 time, doesn't have a name
	func() {
		 fmt.Println("test")
	}()
}