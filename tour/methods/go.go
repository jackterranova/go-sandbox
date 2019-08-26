package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Defining an Abs method on a Vertex type
// Here Vertex v is called a "receiver"
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// Defining an Abs method on a MyFloat type
// (which is just a type def of primitive float64)
//
// You cannot define a method on a built-in type
// The workaround is to create type def on a built-in like below
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Receiver can be a pointer
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale version without ptr type
func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v2 := Vertex{3, 4}
	fmt.Println(v2.Abs())
	v2.Scale(10) // Scale takes a ptr and will change the values of the type
	fmt.Println(v2.Abs())

	// Quirky Go behavior due to the duality of ptr syntax for struct types
	//Above the client code is exactly the same but
	//  the Scale method takes a ptr argument
	// Below the Scale2 method takes a non-ptr Vertex argument
	//   Scale2 operates on a copy of its param
	v3 := Vertex{3, 4}
	fmt.Println(v3.Abs())
	v3.Scale2(10) //
	fmt.Println(v3.Abs())

	// As a function, Scale takes a ptr argument
	// and client code must pass a ptr explicitly
	v4 := Vertex{3, 4}
	Scale(&v4, 10)
	fmt.Println(Abs(v4))
}
