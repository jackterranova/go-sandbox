package main

import "fmt"

// Customary to implement nil handling in interface implementations
// An interface type can be assigned a variable of a declared type only
//  which is a nil object

type I interface {
	M()
	FailBadly()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (t *T) FailBadly() {
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	var t2 *T
	i = t2

	i.FailBadly()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
