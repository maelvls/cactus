package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mo-work/go-technical-test-for-claudia/image"
	"github.com/mo-work/go-technical-test-for-claudia/input"
)

func main() {
	in := input.New(bufio.NewScanner(os.Stdin))
	xAxis, yAxis, err := in.GetImageSize()
	if err != nil {
		fmt.Printf("invalid image value: %s\n", err)
		os.Exit(1)
	}

	bitmap := image.New(xAxis, yAxis)

	commandChan := make(chan input.Command)
	errChan := make(chan error)
	go in.GetEditActions(commandChan, errChan)

	for {
		select {
		case command := <-commandChan:
			switch command.Action {
			case "L":
				bitmap.Set(command.Coords[0], command.Coords[1], command.Char)
			case "S":
				fmt.Println(bitmap.Pretty())
			}

		case err := <-errChan:
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
