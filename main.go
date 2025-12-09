package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/daysix"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(6, daysix.DaySix)
	p.Run()
	fmt.Printf("%+v\n", p)
}
