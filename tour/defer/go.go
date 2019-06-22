package main

import "fmt"

func main() {

	//Can think of defers like a finally in Java.  Useful for things like closing connections
	// and other resource cleanup

	defer fmt.Println("world")

	fmt.Println("hello")
}
