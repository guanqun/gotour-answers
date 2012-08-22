package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	data := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		data[i] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			data[i][j] = uint8((i + j) / 2)
		}
	}

	return data
}

func main() {
	pic.Show(Pic)
}
