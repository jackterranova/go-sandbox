package main

import "fmt"

// An empty interface may hold values of any type.

func main() {
	var i interface{}
	describe(i) // nothing

	// checking that an interface holds a particular type ...
	s, ok := i.(string) //call `s := i.(string)` would result in panic
	fmt.Println(s, ok)

	i = 42
	describe(i) // an int

	i = "hello"
	describe(i) //  a string
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
