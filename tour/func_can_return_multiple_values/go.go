package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func reverse(x, y, z string) (string, string, string) {
	return z, y, x
}

func main() {
	fmt.Println(swap("one", "two"))
	fmt.Println(reverse("one", "two", "three"))
}
