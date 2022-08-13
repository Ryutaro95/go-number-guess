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

	if err := game.receiveNums(); err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(game)
}

func (g *Game) receiveNums() error {
	var str string
	if _, err := fmt.Scan(&str); err != nil {
		return err
	}

	// TODO: asks for input again
	if err := verifyDigit(str); err != nil {
		return err
	}

	for _, s := range str {
		g.playerNums = append(g.playerNums, fmt.Sprintf("%c", s))
	}
	return nil
}

func verifyDigit(nums string) error {
	if len(nums) != digit {
		return errors.New(fmt.Sprintf("%v digit are requried", digit))
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
