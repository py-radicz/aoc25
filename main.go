package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/dayfour"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(4, dayfour.DayFour)
	p.Run()
	fmt.Printf("%+v\n", p)
}
