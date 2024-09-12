package main

import (
	"os"

	"github.com/kish1n/KhOn/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
