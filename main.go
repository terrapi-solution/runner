package main

import (
	"github.com/thomas-illiet/terrapi-runner/cmd"
	"os"
)

func main() {
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
