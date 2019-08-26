package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex //strange syntax:  a map of string->Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m2["Google"])

	var m3 = map[string]Vertex{
		"Facebook": {40.68433, -74.39967},
		"Apple":    {37.42202, -122.08408},
	}

	fmt.Println(m3["Apple"])
}
