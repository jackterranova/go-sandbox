/*

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.

When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

*/

package main

import (
	//"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	println(dx)
	println(dy)
	// allocate the 2-dim array
	// this creates array of size "dy" whose elements are nil-sized arrays
	var retval [][]uint8 = make([][]uint8, dy)

	// iterate over the "dy" sized array and allocate the array elements
	for i, _ := range retval {
		//fmt.Println(j)
		retval[i] = make([]uint8, dx)

		for q, _ := range retval[i] {
			//fmt.Println(r) //should be zero
			//retval[i][q] = uint8(i * q / 2)
			retval[i][q] = uint8((i * q) << uint8(i*q))
		}
	}

	//now we have an initialized 2-dim array (all values are 0)

	//fmt.Println("array", retval)

	return retval
}

func main() {
	pic.Show(Pic)
}
