package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mo-work/go-technical-test-for-claudia/image"
	"github.com/mo-work/go-technical-test-for-claudia/input"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	in := input.New(scanner)
	xAxis, yAxis, err := in.GetImageSize()
	if err != nil {
		fmt.Printf("invalid image value: %s\n", err)
		os.Exit(1)
	}

	bitmap := image.New(xAxis, yAxis)
	//TODO move into input
	for {
		scanner.Scan()
		text := strings.Split(scanner.Text(), " ")
		if text[0] == "S" {
			fmt.Println(bitmap.Pretty())
		}
	}
}
