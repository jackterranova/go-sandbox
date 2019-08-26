package main

import "fmt"

// the fmt package looks for implementation of Stringer to print values

type Person struct {
	Name string
	Age  int
}

// 	implementing Stringer interface

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
