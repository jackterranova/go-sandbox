package main

import "fmt"

//Channels are a typed conduit through which you can send and receive values
// with the channel operator, <-.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c (this will block until the other side is "ready")

	println("sent")
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) //make a channel of type int

	//slice s up to len(s)/2 and send sum of the values to channel c
	println("goroutine 1 starting")
	go sum(s[:len(s)/2], c)

	//slice s the rest of the way and send sum of the values to channel c
	println("goroutine 2 starting")
	go sum(s[len(s)/2:], c)

	//the 2 goroutines above will block until a receive is called

	println("receiving")
	x, y := <-c, <-c // receive from c... the goroutine sends will now unblock

	//By default, sends and receives block until the other side is ready.
	//This allows goroutines to synchronize without explicit locks or condition variables.

	fmt.Println(x, y, x+y)
}
