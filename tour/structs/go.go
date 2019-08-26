package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	mystruct := Vertex{1, 2}

	fmt.Println(mystruct)

	fmt.Println(mystruct.X)
	fmt.Println(mystruct.Y)

	mystruct.X = 1000
	mystruct.Y = 5000

	fmt.Println(mystruct.X)
	fmt.Println(mystruct.Y)

	anotherstruct := Vertex{} //X,Y are both 0
	fmt.Println(anotherstruct)

	anotherstruct = Vertex{Y: 100} //Y 100, X 0
	fmt.Println(anotherstruct)

	var structptr *Vertex = &mystruct

	fmt.Println(structptr)   // <==  Go treats struct pointers differently, no deferencing needed
	fmt.Println(structptr.X) // structptr.X and (*structptr).X are equivalent
	fmt.Println((*structptr).X)

	fmt.Println(&structptr) // <==  If we want the actual pointer value of the struct, we can get its
	//   address even though its already a pointer

}
