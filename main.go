package main

import (
	"os"

	"github.com/thomas-illiet/terrapi-runner/pkg/command"
)

func main() {
	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
