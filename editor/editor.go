package editor

import (
	"errors"
	"strings"
)

type Editor struct {
	Grid [][]string
	rows int
	cols int
}

func New(c, r int) Editor {
	e := Editor{rows: r, cols: c}
	e.clear()
	return e
}

func (e *Editor) Set(x, y int, char string) error {
	if x > e.cols || y > e.rows {
		return errors.New("given coordinate is beyond image grid")
	}

	e.Grid[y-1][x-1] = char

	return nil
}

func (e *Editor) SetMultiY(x, y1, y2 int, char string) error {
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	var err error
	for y := y1; y <= y2; y++ {
		err = e.Set(x, y, char)
	}

	return err
}

func (e *Editor) SetMultiX(x1, x2, y int, char string) error {
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	var err error
	for x := x1; x <= x2; x++ {
		err = e.Set(x, y, char)
	}

	return err
}

func (e *Editor) Pretty() string {
	out := ""
	for x := range e.Grid {
		out += strings.Join(e.Grid[x], "") + "\n"
	}

	return out
}

func (e *Editor) Clear() {
	e.clear()
}

func (e *Editor) clear() {
	grid := make([][]string, e.rows)

	for j := 0; j < e.rows; j++ {
		grid[j] = make([]string, e.cols)
		for k := 0; k < e.cols; k++ {
			grid[j][k] = "O"
		}
	}

	e.Grid = grid
}
