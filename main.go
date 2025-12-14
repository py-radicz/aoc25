package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/daynine"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(daynine.DayNine)
	p.Run()
	fmt.Printf("%+v\n", p)
}
