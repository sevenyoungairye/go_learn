package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func init() {
	m := Image{}
	pic.ShowImage(m)
}

func (m Image) ColorModel() color.Model {

	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {

	return image.Rect(0, 0, 9, 5)
}

func (m Image) At(x, y int) color.Color {

	return color.RGBA{5, 5, 255, 255}
}
