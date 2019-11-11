package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mo-work/go-technical-test-for-claudia/editor"
	"github.com/mo-work/go-technical-test-for-claudia/runner"
)

func main() {
	editor := editor.Editor{}
	r := runner.New(bufio.NewScanner(os.Stdin), os.Stdout, &editor)

	if err := r.ProcessImageSize(); err != nil {
		fmt.Printf("invalid image value: %s\n", err)
	}

	r.ProcessEditActions()

	os.Exit(0)
}
