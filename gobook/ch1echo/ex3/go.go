// measure the time between a string building version and a verion that uses String.Join

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	slowArgs()
	fmt.Println("Took " + strconv.FormatInt(time.Since(start).Nanoseconds(), 10))

	start2 := time.Now()
	lessSlowArgs()
	fmt.Println("Took " + strconv.FormatInt(time.Since(start2).Nanoseconds(), 10))
}

func slowArgs() {
	var s, sep string // default init values for string are ""

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

}

func lessSlowArgs() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
