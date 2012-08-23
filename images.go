package main

import (
	"image"
	"image/color"
	"code.google.com/p/go-tour/pic"
)

type Image struct{
	width int
	height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y ) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{10, 10}
	// it passes 'm' directly so we need to use (i Image) in above interface declaration.
	pic.ShowImage(m)
}
