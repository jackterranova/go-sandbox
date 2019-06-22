package main

import "fmt"

func main() {
	var i, j int = 1, 2
	var x, y = 3, 4
	k := 5 // can only use := in a function
	// can omit var and type
	c, python, java := true, false, "no!"

	fmt.Println(i, j, x, y, k, c, python, java)
}
