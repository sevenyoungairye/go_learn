package main

import (
	"golang.org/x/tour/pic"
)

// the document:.. https://pkg.go.dev/golang.org/x/tour/pic#example-Show
// data:image/png;base64, xxxxxxxxx
func Pic(dx, dy int) [][]uint8 {

	ss := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// x * y
			s[x] = uint8((x + y) / 2)
		}
		ss[y] = s
	}
	return ss
}

func init() {
	pic.Show(Pic)
}
