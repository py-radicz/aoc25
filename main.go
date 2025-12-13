package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/dayeight"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(dayeight.DayEight)
	p.Run()
	fmt.Printf("%+v\n", p)
}
