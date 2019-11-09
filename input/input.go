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

	xAxis, yAxis, err := translateInts(text[1], text[2])
	if err != nil {
		return 0, 0, err
	}

	if !valid(xAxis) || !valid(yAxis) {
		return 0, 0, fmt.Errorf("image axis out of range: %d <= M,N <= %d", MinValue, MaxValue)
	}

	return xAxis, yAxis, nil
}

func translateInts(xAxStr, yAxStr string) (int, int, error) {
	x, err := strconv.Atoi(xAxStr)
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse non-integer '%s'", xAxStr)
	}

	y, err := strconv.Atoi(yAxStr)
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse non-integer '%s'", yAxStr)
	}

	return x, y, nil
}

func valid(axis int) bool {
	return axis > MinValue && axis < MaxValue
}
