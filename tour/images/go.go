package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type MyImage struct {
	width  int
	height int
}

func (myimage MyImage) ColorModel() color.Model {
	return color.Alpha16Model
}

func (myimage MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, myimage.width, myimage.height)
}

func (img MyImage) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		return uint8(x * y)
		//return uint8((x + y) / 2)
		//return uint8(x ^ y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := MyImage{256, 64}
	pic.ShowImage(m)
}
