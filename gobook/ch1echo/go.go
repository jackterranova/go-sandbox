package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // default init values for string are ""

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
