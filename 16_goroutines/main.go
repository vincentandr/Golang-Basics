package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // synchronizes multiple goroutines together (in this case main and goroutine on line 13). To wait for group of goroutines to finish
var counter = 0
var m = sync.RWMutex{} // to protect data access

func main() {
	// go run --race <filename> detects if there is a race condition such as this example. Very useful when coding concurrency
	// msg := "Hello"
	// wg.Add(1) // Add 1 group
	// go func(msg string){
	// 	fmt.Println(msg)
	// 	wg.Done() // Decrement 1 group
	// }(msg) // this copies the msg var from outer function (main) to the goroutine, so changing the msg on line 14 wont change the one inside the goroutine
	// msg = "Goodbye"
	// wg.Wait() // Check if there is 1 or more groups still not done yet, instead of using time to artifically delay main function finish.

	// #2
	// for i := 0; i < 10; i ++ {
	// 	// This works, but locking and unlocking constantly makes it like a sequential execution (no benefit to concurrency)
	// 	wg.Add(2)
	// 	m.RLock()
	// 	go sayHello()
	// 	m.Lock()
	// 	go increment()
	// }
	// wg.Wait()

	// #3
	runtime.GOMAXPROCS(1) // running 1 thread (parameter is no. of threads), if parameter -1 then just print no. of threads available
	
}

// #2
func sayHello(){
	// Lock outside the goroutine, don't lock here because if lock here the goroutines can still be spawned while haven't unlocked yet
	fmt.Printf("Hello #%d\n", counter)
	m.RUnlock()
	wg.Done()
}

// #2
func increment(){
	counter++
	m.Unlock()
	wg.Done()
}