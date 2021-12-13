package main

import "fmt"

func main () {
	ids := []int{1,5,8,10,14}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	// While loop
	for sum < 10 {
		fmt.Println(sum)
		sum++
	}

	emails := map[string] string {"Bob":"bob@gmail.com", "Vincent":"vincent@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
}