package main

import (
	"fmt"
	"io"
	"strings"
)

// Go has many built-in implementation of stream readers,
// including files, networks, connections and string ...

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { // much like C we manually look for EOF to know the stream is done
			break
		}
	}
}
