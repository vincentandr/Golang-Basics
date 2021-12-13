package main

// for printing
import "fmt"

func main() {
	var str string = "This is a string"
	var num = 33
	var flag = true
	const constant = "cannot be changed"

	// alternative to creating a var
	name := "vincent"
	email, address := "email@gmail.com", "address here"

	fmt.Println(str, num, flag, name, constant)
	fmt.Println(email, address)
	fmt.Printf("%T\n", num) // %T find the type of a variable
}