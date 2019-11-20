package runner

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	MinValue = 1
	MaxValue = 1024
)

type Runner struct {
	scanner *bufio.Scanner
	out     io.Writer
	editor  ImageEditor
}

type Command struct {
	Action string
	Coords []int
	Char   string
}

//go:generate counterfeiter . ImageEditor
type ImageEditor interface {
	CreateImage(rows, cols int)
	Set(x, y int, char string) error
	SetMultiY(x, y1, y2 int, char string) error
	SetMultiX(x1, x2, y int, char string) error
	Pretty() string
	Clear()
}

func New(reader *bufio.Scanner, writer io.Writer, ed ImageEditor) Runner {
	return Runner{scanner: reader, out: writer, editor: ed}
}

func (r Runner) ProcessImageSize() error {
	r.scanner.Scan()
	text := strings.Split(r.scanner.Text(), " ")

	if strings.ToUpper(text[0]) != "I" {
		return fmt.Errorf("unrecognised command '%s', use 'I' for Image initialisation", text[0])
	}

	axes, err := translateInts(text[1:])
	if err != nil {
		return err
	}
	xAxis, yAxis := axes[0], axes[1]

	if !valid(xAxis) || !valid(yAxis) {
		return fmt.Errorf("image axis out of range: %d <= M,N <= %d", MinValue, MaxValue)
	}

	r.editor.CreateImage(xAxis, yAxis)

	return nil
}

func (r Runner) ProcessEditActions() {
	for {
		r.scanner.Scan()
		text := strings.Split(r.scanner.Text(), " ")

		if text[0] == "" {
			break
		}

		command := Command{Action: strings.ToUpper(text[0])}

		if len(text) > 1 {
			coords, err := translateInts(text[1 : len(text)-1])
			if err != nil {
				fmt.Fprintln(r.out, err)
				continue
			}
			command.Coords = coords
			command.Char = strings.ToUpper(text[len(text)-1])
		}

		if err := r.applyAction(command); err != nil {
			fmt.Fprintln(r.out, err)
			continue
		}
	}
}

func (r Runner) applyAction(command Command) error {
	switch command.Action {
	case "L":
		if err := r.editor.Set(command.Coords[0], command.Coords[1], command.Char); err != nil {
			return err
		}
	case "V":
		if err := r.editor.SetMultiY(command.Coords[0], command.Coords[1], command.Coords[2], command.Char); err != nil {
			return err
		}
	case "H":
		if err := r.editor.SetMultiX(command.Coords[0], command.Coords[1], command.Coords[2], command.Char); err != nil {
			return err
		}
	case "S":
		fmt.Fprintln(r.out, r.editor.Pretty())
	case "C":
		r.editor.Clear()
	default:
		fmt.Fprintln(r.out, "invalid action")
	}

	return nil
}

func translateInts(axStr []string) ([]int, error) {
	axes := []int{}

	for _, a := range axStr {
		axis, err := strconv.Atoi(a)
		if err != nil {
			return nil, fmt.Errorf("could not parse non-integer '%s'", a)
		}
		axes = append(axes, axis)
	}

	return axes, nil
}

func valid(axis int) bool {
	return axis > MinValue && axis < MaxValue
}
