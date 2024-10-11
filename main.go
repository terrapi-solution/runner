package main

import (
	"github.com/thomas-illiet/terrapi-runner/command"
	"os"
)

func main() {
	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
