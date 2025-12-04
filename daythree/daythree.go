/* Package daythree */
package daythree

import (
	"log"
	"strings"

	"github.com/py-radicz/aoc25/utils"
)

func MaxJoltage(s string, k int) string {
	if len(s) == 0 || k == 0 || k > len(s) {
		return ""
	}

	maxStart := len(s) - k
	bestIdx := 0

	for i := 1; i <= maxStart; i++ {
		if s[i] > s[bestIdx] {
			bestIdx = i
		}
	}
	return string(s[bestIdx]) + MaxJoltage(s[bestIdx+1:], k-1)
}

func DayThree(day int) (partOne, partTwo int) {
	// in := `987654321111111
	// 811111111111119
	// 234234234234278
	// 818181911112111`
	in, err := utils.GetInput(day)
	if err != nil {
		log.Fatalf("failed to load Input for day: %d, %s", day, err)
	}

	// part one
	for _, b := range strings.Fields(string(in)) {
		mj := MaxJoltage(b, 2)
		num := 0
		for _, ch := range mj {
			num = num*10 + int(ch-'0')
		}
		partOne += num
	}

	// part two
	for _, b := range strings.Fields(string(in)) {
		mj := MaxJoltage(b, 12)
		num := 0
		for _, ch := range mj {
			num = num*10 + int(ch-'0')
		}
		partTwo += num
	}

	return
}
