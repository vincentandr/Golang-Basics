package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Defer skips and only executes in LIFO order (stack) after it reaches the end of function but before returning
	fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("fourth")
	fmt.Println("fifth")

	// Defer can be used to close a resource, e.g. when calling an API and received a response
	a := 10
	defer fmt.Println(a) // Print 10 not 5
	a = 5

	// Panic when program runs into an exception that the program can't recover from
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	panicker()

	fmt.Println("finish") // This line is still executed if panicker() encountered a panic and recovered
}

func panicker() {
	// Recover is used to recover from panic. It only stops executing instructions in current call stack, and will let the program continue to execute in higher call stack.
	defer func(){
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()

	// Panic practical example
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// This might throw an error if the port is blocked / already used. Try opening 2 terminals and both try to use 8080 port
	err := http.ListenAndServe(":8080", nil)
	if err != nil{ 
		panic(err.Error())
	}

	fmt.Println("panicker done") // This will not execute if there's panic
}