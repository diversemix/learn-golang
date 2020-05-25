package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (Image) At(x, y int) color.Color {
	if (x+y)%11 == 0 {
		return color.Black
	} else {
		return color.White
	}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
