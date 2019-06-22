package main

import (
	"fmt"
	"math"
)

// if statements can be with or without parens
// but braces are always required

func sqrt2(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// short statemtns in an if... v is scoped to the if block
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(sqrt2(2), sqrt2(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
