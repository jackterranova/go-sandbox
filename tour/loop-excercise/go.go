package main

/*
As a way to play with functions and loops, let's implement a square root function: given a number x,
we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop.

Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

*/

import (
	"fmt"
)

func Sqrt_unnamed_return(x float64) float64 {
	return 100
}

func Sqrt_named_return(x float64) (retval float64) {
	retval = 100
	return
}

func Sqrt_named_return_but_not_set(x float64) (retval float64) {
	return ///this returns 0!
}

func Sqrt_div_by_zero(x float64) (retval float64) {

	retval -= (retval*retval - x) / (2 * retval) // << This does not blowup!!!

	return
}

func Sqrt_10times_method(x float64) (retval float64) {
	retval = 1.0

	for i := 0; i < 10; i++ {
		retval -= (retval*retval - x) / (2 * retval) // << This does not blowup!!!
		fmt.Println(retval)
	}

	return
}

func Sqrt_until_close(x float64) (retval float64) {
	retval = 1.0
	i := 0

	// stop once we are within .05
	for ; x-retval > .05 && i < 10000; i++ {

		retval -= (retval*retval - x) / (2 * retval) // << This does not blowup!!!
		fmt.Println(i, retval)
	}

	fmt.Println("final loop count ", i)

	return
}

func main() {
	fmt.Println(Sqrt_unnamed_return(2))
	fmt.Println(Sqrt_named_return(2))
	fmt.Println(Sqrt_named_return_but_not_set(2))
	fmt.Println(Sqrt_div_by_zero(2))
	fmt.Println(Sqrt_10times_method(2))
	fmt.Println(Sqrt_until_close(2))
}
