// modify to print the name of the program

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // default init values for string are ""

	if len(os.Args) > 0 {
		fmt.Println("Program name ==> " + os.Args[0])
	}

	fmt.Print("Params => ex1")

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
