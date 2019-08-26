package main

import "fmt"

func main() {
	// use make to dynamically create slices

	a := make([]int, 5) // size 5 with zero'ed elements
	printSlice("a", a)

	b := make([]int, 0, 5) // size 0, with a capacity of 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
