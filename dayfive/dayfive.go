/* Package dayfive */
package dayfive

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/py-radicz/aoc25/utils"
)

const day int = 5

type Range struct {
	first int
	last  int
}

func (r *Range) Contains(ingr int) bool {
	return ingr >= r.first && ingr <= r.last
}

func (r *Range) Count() int {
	return (r.last - r.first) + 1
}

func mergeRanges(ranges []Range) []Range {
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.first - b.first
	})
	merged := []Range{ranges[0]}

	for _, rg := range ranges[1:] {
		last := &merged[len(merged)-1]
		if rg.first <= last.last+1 {
			if rg.last > last.last {
				last.last = rg.last
			}
		} else {
			merged = append(merged, rg)
		}
	}
	return merged
}

func ParseInput(in string) (ranges []Range, ingredients []int) {
	for _, row := range strings.Fields(in) {
		if strings.Contains(row, "-") {
			rg := strings.Split(row, "-")
			first, err := strconv.Atoi(rg[0])
			if err != nil {
				log.Fatal("failed to parse range")
			}

			last, err := strconv.Atoi(rg[1])
			if err != nil {
				log.Fatal("failed to parse range")
			}
			ranges = append(ranges, Range{first: first, last: last})
			continue
		}

		if len(strings.Fields(row)) == 0 {
			continue
		}

		num, err := strconv.Atoi(row)
		if err != nil {
			log.Fatal("failed to parse ingredients")
		}
		ingredients = append(ingredients, num)
	}

	return
}

func DayFive() (partOne, partTwo int) {
	//in := `3-5
	//10-14
	//16-20
	//12-18
	//
	//1
	//5
	//8
	//11
	//17
	//32`

	input, err := utils.GetInput(day)
	if err != nil {
		log.Fatal("failed to load input for day", day)
	}
	in := string(input)

	// part one
	ranges, ingredients := ParseInput(in)

	for _, ing := range ingredients {
		for _, rg := range ranges {
			if rg.Contains(ing) {
				partOne++
				break
			}
		}
	}

	// part two
	for _, rg := range mergeRanges(ranges) {
		partTwo += rg.Count()
	}
	return
}
