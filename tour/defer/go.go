package main

import "fmt"

func main() {

	//Can think of defers like a finally in Java.  Useful for things like closing connections
	// and other resource cleanup

	//defer statements are placed in stack and hence are executed in LIFO order

	defer fmt.Println("world3") //prints last
	defer fmt.Println("world2")
	defer fmt.Println("world1") //prints first

	fmt.Println("hello")
}
