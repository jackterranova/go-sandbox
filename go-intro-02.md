## Building a Go Library

1. mkdir $GOPATH/src/github.com/jackterranova/go-sandbox/stringutil

2.  Create `reverse.go` in the new directory ...

```
// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
``` 

3. `go build github.com/jackterranova/go-sandbox/stringutil`

__Note this doesn't produce an output file.__  The binary is stored in a build cache.  

Also note the use of `build` (since we're building a library) as opposed to `install` (of an executable).

4.  Modify `hello.go` ...

Note the package `stringutil`, despite the fact that the file's path is $GOPATH/github.com/jackterranova/go-sandbox/stringutil.  

> Go package naming conventions suggest the last element of the path be used as the package name.  
> However, the import string (as used in hello.go) is the full path under $GOPATH

```
package main

import (
	"fmt"

	"github.com/jackterranova/go-sandbox/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}
```

5. go install github.com/jackterranova/go-sandbox/hello

6. `$GOPATH/bin/hello` or `hello` (again, if $GOPATH/bin is your executable path)