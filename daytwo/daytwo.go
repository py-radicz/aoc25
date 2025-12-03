/* Package daytwo */
package daytwo

import (
	"bytes"
	"log"
	"slices"
	"strconv"

	"github.com/py-radicz/aoc25/utils"
)

type Range struct {
	first int
	last  int
}

func IsInvalid(id string, partition int) bool {
	parts := make(map[string]struct{}, partition)

	for chunk := range slices.Chunk([]byte(id), partition) {
		parts[string(chunk)] = struct{}{}
	}
	return len(parts) == 1
}

func GetRanges(input []byte) (res []Range) {
	for _, rg := range bytes.Split(input, []byte(",")) {
		a := bytes.Split(rg, []byte("-"))
		res = append(res, Range{first: utils.AtoiBytes(a[0]), last: utils.AtoiBytes(a[1])})
	}
	return
}

func DayTwo(day int) (partOne, partTwo int) {
	// in := []byte(`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`)
	in, err := utils.GetInput(day)
	if err != nil {
		log.Fatalf("failed to load Input for day: %d, %s", day, err)
	}

	for _, rg := range GetRanges(in) {
		for i := rg.first; i <= rg.last; i++ {
			idstr := strconv.Itoa(i)

			// Part One
			// only id with even number of digits can make Invalid ID
			if len(idstr)%2 == 0 {
				if IsInvalid(idstr, len(idstr)/2) {
					partOne += i
				}
			}

			// Part Two
			for j := 1; j <= len(idstr)/2; j++ {
				if IsInvalid(idstr, j) {
					partTwo += i
					break
				}
			}
		}
	}
	return partOne, partTwo
}
