package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/daythree"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(3, daythree.DayThree)
	p.Run()
	fmt.Printf("%+v\n", p)
}
