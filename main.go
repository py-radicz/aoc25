package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/dayfive"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(5, dayfive.DayFive)
	p.Run()
	fmt.Printf("%+v\n", p)
}
