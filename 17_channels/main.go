package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // Integer 50 means create a channel that has an internal data store with 50 integers store

	// wg.Add(2)
	// go func() {
	// 	i := <- ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// go func() {
	// 	i := 42
	// 	ch <- i
	// 	i = 27
	// 	wg.Done()
	// }()

	// for j := 0; j < 5; j++ {
	// 	wg.Add(2)
	// 		go func() {
	// 			i := <- ch
	// 			fmt.Println(i)
	// 			wg.Done()
	// 		}()
	// 		go func() {
	// 			i := 42
	// 			ch <- i
	// 			i = 27
	// 			wg.Done()
	// 		}()
	// }

	// wg.Add(2)
	// go func(ch <-chan int) { // Restricts this thread to only receive channel
	// 	i := <- ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) { // Send only channel. Similar to polymorphic where the channel is cast as a send channel
	// 	i := 42
	// 	ch <- i
	// 	i = 27
	// 	wg.Done()
	// }(ch)

	// wg.Add(2)
	// go func(ch <-chan int) { // Restricts this thread to only receive channel
	// 	for val := range ch {
	// 		fmt.Println(val)
	// 	}
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) { // Send only channel. Similar to polymorphic where the channel is cast as a send channel
	// 	ch <- 42
	// 	ch <- 27 // Send and receive must occur the same amount of times
	// 	close(ch) // To close the channel indicating the receiving channel there is no more value to be sent
	// 	wg.Done()
	// }(ch)

	// wg.Add(2)
	// go func(ch <-chan int) { // Restricts this thread to only receive channel
	// 	for {
	// 		if val, ok := <- ch; ok {
	// 			fmt.Println(val)
	// 		} else { // To know that channel is closed
	// 			break
	// 		}
	// 	}
	// 	wg.Done()
	// }(ch)
	// go func(ch chan<- int) { // Send only channel. Similar to polymorphic where the channel is cast as a send channel
	// 	ch <- 42
	// 	ch <- 27 // Send and receive must occur the same amount of times
	// 	close(ch) // To close the channel indicating the receiving channel there is no more value to be sent
	// 	wg.Done()
	// }(ch)

	doneCh := make(chan struct{})
	go func(ch <-chan int) { // Restricts this thread to only receive channel
		for {
			select { // Blocking until a message comes in. Also allows to listen to more than 1 channel
				case val := <- ch:
					fmt.Println(val)
				case <- doneCh:
					break
				// default: // This makes select becomes non-blocking. It will go here if there's no msg coming in
			}
		}
	}(ch)
	go func(ch chan<- int) { // Send only channel. Similar to polymorphic where the channel is cast as a send channel
		ch <- 42
		ch <- 27 // Send and receive must occur the same amount of times
	}(ch)
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} // To close ch channel. struct{} costs 0 memory alloc if there's no field.
}