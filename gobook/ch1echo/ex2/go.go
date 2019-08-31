// modify to print the index and value of each param on a separate line

package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 0 {
		fmt.Println("Program name ==> " + os.Args[0])
	}

	fmt.Print("Params => ...")

	for i, arg := range os.Args {
		fmt.Println(string(i) + arg)
	}
}
