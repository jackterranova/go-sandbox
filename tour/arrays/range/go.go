package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	// range returns 2 values: index and copy of the value at the index
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// omit returned values that you don't care about
	for i, _ := range pow {
		fmt.Printf("index %d\n", i)
	}

	// omit returned values that you don't care about
	for _, v := range pow {
		fmt.Printf("values %d\n", v)
	}
}
