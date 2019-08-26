package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slicing primes
	fmt.Println(primes[4:5]) // the right hand limit is NOT inclusive
	fmt.Println(primes[4:6]) // the right hand limit is NOT inclusive
	fmt.Println(primes[5])   // this can be confusing given arrays are 0-based
	fmt.Println(primes[:3])  // the first 3
	fmt.Println(primes[3:])  // the last 3
	fmt.Println(primes[:])   // the whole array

	//Slices are pointers to the underlying array
	//changing through a slice changes the underlying array
	slice := primes[4:6]
	slice[0] = 17
	slice[1] = 19
	fmt.Println(primes)

	// len and cap
	fmt.Println(len(primes))
	fmt.Println(cap(primes)) // -len and cap of an array are always the same
	fmt.Println(len(slice))  // -len of a slice is the number of elements in the slice
	fmt.Println(cap(slice))  // -cap of a slice is size of the array it points
	//  to starting at the first position of the slice

	slice = primes[1:3] // reslice for the 2nd/3rd element
	fmt.Println(slice)
	fmt.Println(cap(slice)) // now cap is 5 because the slice can be extened
	// past its high end value

}
