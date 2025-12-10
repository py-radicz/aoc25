package main

import (
	"fmt"

	"github.com/py-radicz/aoc25/dayseven"
	"github.com/py-radicz/aoc25/utils"
)

func main() {
	p := utils.NewPuzzle(7, dayseven.DaySeven)
	p.Run()
	fmt.Printf("%+v\n", p)
}
