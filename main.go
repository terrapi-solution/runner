package main

import (
	"github.com/terrapi-solution/runner/cmd"
	"os"
)

func main() {
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
