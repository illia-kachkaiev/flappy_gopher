package main

import (
	"fmt"
	"os"
)

const (
	gravity      = 9.8 // meters per square second
	windowHeight = 600 // pixels
	windowWidth  = 800 // pixels
)

func main() {
	game := game{}

	if err := game.run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}
