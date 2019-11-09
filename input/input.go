package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	MinValue = 1
	MaxValue = 1024
)

type Input struct {
	scanner *bufio.Scanner
}

type Command struct {
	Action string
	Coords []int
	Char   string
}

func New(reader *bufio.Scanner) Input {
	return Input{scanner: reader}
}

func (i *Input) GetImageSize() (int, int, error) {
	i.scanner.Scan()
	text := strings.Split(i.scanner.Text(), " ")

	if text[0] != "I" {
		return 0, 0, fmt.Errorf("unrecognised command '%s', use 'I' for Image initialisation", text[0])
	}

	axes, err := translateInts(text[1:])
	if err != nil {
		return 0, 0, err
	}
	xAxis, yAxis := axes[0], axes[1]

	if !valid(xAxis) || !valid(yAxis) {
		return 0, 0, fmt.Errorf("image axis out of range: %d <= M,N <= %d", MinValue, MaxValue)
	}

	return xAxis, yAxis, nil
}

func (i *Input) GetEditActions(actionChan chan Command, errChan chan error) {
	for {
		i.scanner.Scan()
		text := strings.Split(i.scanner.Text(), " ")

		command := Command{Action: text[0]}

		if len(text) > 1 {
			coords, err := translateInts(text[1 : len(text)-1])
			if err != nil {
				errChan <- err
				continue
			}
			command.Coords = coords
			command.Char = text[len(text)-1]
		}

		actionChan <- command
	}
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
