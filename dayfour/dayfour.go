/* Package dayfour */
package dayfour

import (
	"log"
	"strings"

	"github.com/py-radicz/aoc25/utils"
)

func EnhanceInput(lines []string) [][]rune {
	// adds border so I am never out of index
	cols := len(lines[0])
	border := []rune(strings.Repeat(".", cols+2))

	enhanced := make([][]rune, 0, len(lines)+2)
	enhanced = append(enhanced, border)

	for _, line := range lines {
		row := make([]rune, 0, cols+2)
		row = append(row, '.')
		row = append(row, []rune(line)...)
		row = append(row, '.')
		enhanced = append(enhanced, row)
	}

	enhanced = append(enhanced, border)
	return enhanced
}

type Position struct {
	x, y int
}

func RemoveRolls(grid [][]rune) int {
	var removed []Position
	for i, r := range grid {
		for j := range r {
			if grid[i][j] != '@' {
				continue
			}

			if i != 0 && j != 0 && i != len(grid)-1 && j != len(r)-1 {
				neighbours := string(grid[i-1][j-1:j+2]) + string(grid[i+1][j-1:j+2]) + string(grid[i][j-1]) + string(grid[i][j+1])
				if strings.Count(neighbours, "@") < 4 {
					removed = append(removed, Position{i, j})
				}
			}

		}
	}

	for _, pos := range removed {
		grid[pos.x][pos.y] = '.'
	}
	return len(removed)
}

func DayFour(day int) (partOne, partTwo int) {
	// in := `..@@.@@@@.
	//    @@@.@.@.@@
	//    @@@@@.@.@@
	//    @.@@@@..@.
	//    @@.@@@@.@@
	//    .@@@@@@@.@
	//    .@.@.@.@@@
	//    @.@@@.@@@@
	//    .@@@@@@@@.
	//    @.@.@@@.@.`

	input, err := utils.GetInput(day)
	in := string(input)
	if err != nil {
		log.Fatal("failed to load input")
	}
	grid := EnhanceInput(strings.Fields(in))

	for i := 0; ; i++ {
		removed := RemoveRolls(grid)
		if removed == 0 {
			break
		}

		if i == 0 {
			partOne = removed
		}

		partTwo += removed

	}

	return
}
