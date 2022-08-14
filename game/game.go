package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Digit       int
	PlayerNums  []string
	CorrectNums []string
}

func New() Game {
	var game Game
	return game
}

func (g *Game) Start() error {
	g.generateNums()
	for {
		str, err := g.receiveNums()
		if err != nil {
			return err
		}

		if err := g.verifyDigit(str); err != nil {
			fmt.Println(err)
			continue
		}

		match, include := g.hoge()
		if match == 3 {
			break
		} else {
			fmt.Printf("Matched number: %v, Included number: %v\n", match, include)
			continue
		}
	}
	fmt.Println("Congratiration!")
	return nil
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
		return fmt.Errorf("%v digit are required", g.Digit)
	}
	return nil
}

func (g *Game) generateNums() {
	rand.Seed(time.Now().UnixNano())
	list := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := 0; i < g.Digit; i++ {
		randomInt := rand.Intn(len(list))
		g.CorrectNums = append(g.CorrectNums, fmt.Sprint(list[randomInt]))
		list = append(list[:randomInt], list[randomInt+1:]...)
	}
}

func (g *Game) hoge() (int, int) {
	// fmt.Println(g.PlayerNums, g.CorrectNums)
	var includeNum int
	var matchNum int
	for i := 0; i < g.Digit; i++ {
		if g.PlayerNums[i] == g.CorrectNums[i] {
			matchNum++
		} else if contains(g.PlayerNums[i], g.CorrectNums) {
			includeNum++
		}
	}

	return matchNum, includeNum
}

func contains(a string, b []string) bool {
	for _, c := range b {
		if c == a {
			return true
		}
	}
	return false
}
