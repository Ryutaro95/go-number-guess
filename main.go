package main

import (
	"fmt"
	"os"

	"github.com/Ryutaro95/go-number-guess/game"
)

const digit = 3

func main() {
	game := game.New()
	game.Digit = digit
	if err := game.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
}
