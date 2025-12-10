/* Package dayseven */
package dayseven

import (
	"bytes"
	"log"
	"slices"

	"github.com/py-radicz/aoc25/utils"
)

//var in []byte = []byte(`.......S.......
//...............
//.......^.......
//...............
//......^.^......
//...............
//.....^.^.^.....
//...............
//....^.^...^....
//...............
//...^.^...^.^...
//...............
//..^...^.....^..
//...............
//.^.^.^.^.^...^.
//...............`)

func DaySeven(day int) (partOne, partTwo int) {
	in, err := utils.GetInput(day)
	if err != nil {
		log.Fatalf("failed to load input for day %d", day)
	}

	matrix := bytes.Split(in, []byte("\n"))

	// part one
	for i, row := range matrix {
		if i == 0 {
			continue
		}
		for j, col := range row {
			if (matrix[i-1][j] == '|' || matrix[i-1][j] == 'S') && col == '.' {
				matrix[i][j] = '|'
			}
			if matrix[i-1][j] == '|' && col == '^' {
				partOne++
				matrix[i][j-1] = '|'
				matrix[i][j+1] = '|'
			}
		}
	}

	// part two
	paths := make([][]int, len(matrix))
	cols, rows := len(matrix[0]), len(matrix)-1

	for i := range paths {
		paths[i] = make([]int, cols)
	}

	// set number of paths on each endpoint
	for j := range cols {
		if matrix[rows][j] == '|' {
			paths[rows][j] = 1
		}
	}

	for i := rows - 1; i > 0; i-- {
		for j := range cols {
			// pass number of paths further the beam
			if matrix[i][j] == '|' && matrix[i+1][j] == '|' {
				paths[i][j] = paths[i+1][j]
			}

			// sum number of paths on splitter
			if matrix[i][j] == '|' && matrix[i+1][j] == '^' {
				paths[i][j] = paths[i+1][j-1] + paths[i+1][j+1]
			}
		}
	}

	partTwo = slices.Max(paths[1])
	return
}
