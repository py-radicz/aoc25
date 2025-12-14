/* Package dayeight */
package dayeight

import (
	"log"
	"maps"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/py-radicz/aoc25/utils"
)

const day int = 8

var in string = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

type Point struct {
	x, y, z int
	circuit int
}

type Pair struct {
	p1       *Point
	p2       *Point
	distance int
}

func (p *Pair) SetDistance() {
	dx, dy, dz := p.p1.x-p.p2.x, p.p1.y-p.p2.y, p.p1.z-p.p2.z
	// sqrt does not have ordering impact
	p.distance = (dx * dx) + (dy * dy) + (dz * dz)
}

func GetClosestPairs(points []*Point) []Pair {
	pairs := make([]Pair, 0, (len(points)*len(points)-1)/2)
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			pair := Pair{p1: p, p2: points[j]}
			pair.SetDistance()
			pairs = append(pairs, pair)
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance < pairs[j].distance
	})
	return pairs
}

func GetPoints(input string) []*Point {
	rows := strings.Fields(input)
	points := make([]*Point, len(rows))

	for i, row := range rows {
		coords := strings.Split(row, ",")
		nums := [3]int{}
		for i, coord := range coords {
			num, err := strconv.Atoi(coord)
			if err != nil {
				log.Fatal("failed to parse input")
			}
			nums[i] = num
		}
		points[i] = &Point{x: nums[0], y: nums[1], z: nums[2]}
	}

	return points
}

func ConnectPairs(in string, maxPairs int) int {
	points := GetPoints(in)
	pairs := GetClosestPairs(points)

	circuitID := 1

	for i := range len(pairs) {
		// partOne
		if i == maxPairs {
			results := make(map[int]int)
			for _, p := range points {
				if p.circuit == 0 {
					continue
				}
				results[p.circuit]++
			}
			circuitLens := slices.Collect(maps.Values(results))
			sort.Slice(circuitLens, func(i, j int) bool {
				return circuitLens[i] > circuitLens[j]
			})

			return circuitLens[0] * circuitLens[1] * circuitLens[2]
		}

		// part Two
		if i != 0 && CircuitFullyMerged(points) {
			return pairs[i-1].p1.x * pairs[i-1].p2.x
		}

		pair := pairs[i]

		// brand new circuit
		if pair.p1.circuit == 0 && pair.p2.circuit == 0 {
			pair.p1.circuit, pair.p2.circuit = circuitID, circuitID
			circuitID++
			continue
		}
		// circuit exists, add member
		if pair.p1.circuit == 0 || pair.p2.circuit == 0 {
			pair.p1.circuit |= pair.p2.circuit
			pair.p2.circuit |= pair.p1.circuit
			continue
		}

		// two different circuits sharing connection, merge into one
		if pair.p1.circuit != pair.p2.circuit {
			newCircuitID := pair.p1.circuit
			oldCircuitID := pair.p2.circuit

			for _, point := range points {
				if point.circuit == oldCircuitID {
					point.circuit = newCircuitID
				}
			}

		}
	}
	return 0
}

func CircuitFullyMerged(points []*Point) bool {
	merged := true
	c := points[0].circuit
	for _, p := range points[1:] {
		if p.circuit != c {
			merged = false
		}
	}
	return merged
}

func DayEight() (partOne, partTwo int) {
	input, err := utils.GetInput(day)
	if err != nil {
		log.Fatalf("failed to load input for day %d", day)
	}
	in := string(input)

	partOne = ConnectPairs(in, 1000)
	partTwo = ConnectPairs(in, -1)

	return
}
