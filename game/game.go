package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Digit      int
	PlayerNums []string
	CorrectNum []string
}

func New() Game {
	var game Game
	return game
}

func (g *Game) Start() error {
	for {
		str, err := g.receiveNums()
		if err != nil {
			return err
		}

		if err := g.verifyDigit(str); err != nil {
			fmt.Println(err)
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
		g.PlayerNums = append(g.PlayerNums, fmt.Sprintf("%c", s))
	}
	return str, nil
}

func (g *Game) verifyDigit(nums string) error {
	if len(nums) != g.Digit {
		return fmt.Errorf("%s digit are required", g.Digit)
	}
	return nil
}

func (g *Game) generateNums() {
	rand.Seed(time.Now().UnixNano())
	list := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := 0; i < g.Digit; i++ {
		randomInt := rand.Intn(len(list))
		g.CorrectNum = append(g.CorrectNum, fmt.Sprint(list[randomInt]))
		list = append(list[:randomInt], list[randomInt+1:]...)
	}
}
