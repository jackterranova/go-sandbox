package main

import "fmt"

var i, j int = 1, 2 //<= explicit type
var x, y = 3, 4     //<= type optional with initializer

var (
	k, l, m = 100, 200, 300
)

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, x, y, c, python, java)
	fmt.Println(k, l, m)
}
