package main

import "fmt"

func main() {
	// can't use parens with for loops

	//note the short expression with i ... i is scoped to the for block
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// ";" is optional
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	// compiler will not warn you about this!!!!
	/*
		for {
		}
	*/

}
