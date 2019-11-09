package image

import (
	"strings"
)

type Image struct {
	Grid [][]string
}

func New(c, r int) Image {
	i := Image{Grid: generate(r, c)}
	return i
}

func (i *Image) Set(x, y int, char string) {
	i.Grid[y-1][x-1] = char
}

func (i *Image) Pretty() string {
	out := ""
	for x := range i.Grid {
		out += strings.Join(i.Grid[x], " ") + "\n"
	}

	return out
}

func generate(rows, cols int) [][]string {
	grid := make([][]string, rows)

	for j := 0; j < rows; j++ {
		grid[j] = make([]string, cols)
		for k := 0; k < cols; k++ {
			grid[j][k] = "O"
		}
	}

	return grid
}
