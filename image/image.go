package image

import (
	"strings"
)

type Image struct {
	Grid [][]string
	rows int
	cols int
}

func New(c, r int) Image {
	i := Image{rows: r, cols: c}
	i.clear()
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

func (i *Image) Clear() {
	i.clear()
}

func (i *Image) clear() {
	grid := make([][]string, i.rows)

	for j := 0; j < i.rows; j++ {
		grid[j] = make([]string, i.cols)
		for k := 0; k < i.cols; k++ {
			grid[j][k] = "O"
		}
	}

	i.Grid = grid
}
