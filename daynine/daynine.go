/* Package daynine */
package daynine

import (
	"log"
	"strconv"
	"strings"

	"github.com/py-radicz/aoc25/utils"
)

const day int = 9

var in string = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	x, y int
}

func BiggestArea(points []Point) int {
	biggest := 0

	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			area := (Abs(p1.x-p2.x) + 1) * (Abs(p1.y-p2.y) + 1)
			if area > biggest {
				biggest = area
			}
		}
	}
	return biggest
}

func GetPoints(input string) (res []Point) {
	for _, row := range strings.Fields(input) {
		coords := strings.Split(row, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal("failed to load points")
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal("failed to load points")
		}
		res = append(res, Point{x: x, y: y})

	}
	return res
}

func DayNine() (partOne, partTwo int) {
	input, err := utils.GetInput(day)
	if err != nil {
		log.Fatalf("failed to load input for day %d", day)
	}
	in := string(input)

	points := GetPoints(in)
	partOne = BiggestArea(points)
	return
}
