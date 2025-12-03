package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/daytwo"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(2, daytwo.DayTwo)
	p.Run()
	fmt.Printf("%+v\n", p)
}
