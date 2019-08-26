package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	// switch has no fallthrough
	// 	      can be used on non-int, non-constants like strings
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println()

	// switch can substitute for long if-then-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// using switch to check on a variables type
	var i interface{}
	//i = "test"
	//i = 100
	//i = time.Now()

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case time.Time:
		fmt.Println("Time type")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
