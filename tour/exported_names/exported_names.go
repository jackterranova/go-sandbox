package main

import (
	"fmt"
	"math"
)

func main() {
	// Upper-cased names are exported and visible
	fmt.Println("The value of pi is:", math.Pi)

	// lower-cased names are not exported/visible
	//fmt.Println(math.pi)  <--- this would be a compile error
}
