package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const digit = 3

type Game struct {
	digit      int
	playerNums []string
	correctNum []string
}

func main() {
	var game Game
	game.digit = digit
	if err := game.start(); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
}

func (g *Game) start() error {
	for {
		str, err := g.receiveNums()
		if err != nil {
			return err
		}

		if err := g.verifyDigit(str); err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
	}
}

func (g *Game) receiveNums() (string, error) {
	var str string
	if _, err := fmt.Scan(&str); err != nil {
		return "", err
	}

	for _, s := range str {
		g.playerNums = append(g.playerNums, fmt.Sprintf("%c", s))
	}
	return str, nil
}

func (g *Game) verifyDigit(nums string) error {
	if len(nums) != g.digit {
		return fmt.Errorf("%s digit are required", g.digit)
	}
	return nil
}

func generateNums() []string {
	rand.Seed(time.Now().UnixNano())
	list := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var x []string

	for i := 0; i < digit; i++ {
		randomInt := rand.Intn(len(list))
		x = append(x, fmt.Sprint(list[randomInt]))
		list = append(list[:randomInt], list[randomInt+1:]...)
	}

	return x
}
