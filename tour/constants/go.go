package main

import "fmt"

const Pi = 3.14

const (
	dontExportMe = 999
	ExportMe     = 999.999
)

func main() {
	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	//Truth = false  would result in error

	fmt.Println(dontExportMe)
	fmt.Println(ExportMe)
}
