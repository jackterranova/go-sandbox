package main

import (
	"fmt"
	"strconv"
)

// need only specify the type of consecutive params once
func add_ints(x, y, z int) int {
	return x + y + z
}

func concat(x, y, z string) string {
	return x + y + z
}

// mixed params need expl
func concat2(x int, y string, z int) string {
	return strconv.Itoa(x) + y + strconv.Itoa(z)
}

func main() {
	fmt.Println(add_ints(3, 4, 5))

	fmt.Println(concat("this ", "is", " a test"))

	fmt.Println(concat2(1, " is ", 1))
}
