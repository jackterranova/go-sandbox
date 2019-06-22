package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// -math.rand is package of pseudo-random number generators
	//
	// -these generators are deterministic... they will produce the same number
	// or sequence of numbers each time they are run
	//

	fmt.Println("Random number =>", rand.Int())

	// Seeding with the same number will produce the same results

	fmt.Println("Random number with seed 100 =>", rand.Intn(100))

	// Use a different seed to ensure a different random number

	fmt.Println("Random number with time-based seed =>", rand.Intn(time.Now().Second()))

	fmt.Println(time.Now())
}
