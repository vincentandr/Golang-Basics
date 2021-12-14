package main

import "fmt"

func main() {
	// emails := make(map[string] string)

	// emails["Bob"] = "bob@gmail.com"
	// emails["Vincent"] = "vincent@gmail.com"

	// Map is also a pointer like slices.
	emails := map[string] string {"Bob":"bob@gmail.com", "Vincent":"vincent@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)
}