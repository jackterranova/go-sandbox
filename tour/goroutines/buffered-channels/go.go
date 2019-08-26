package main

import "fmt"

//Channels can be buffered.

func main() {

	ch := make(chan int, 2) //channel buffer size of 2

	//Sends to a buffered channel block only when the buffer is full.
	ch <- 1
	ch <- 2

	//ch <- 3  //overflowing the buffer will cause a deadlock error

	//Receives block when the buffer is empty.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
