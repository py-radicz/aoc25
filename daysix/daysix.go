/* Pacakge daysix */
package daysix

import (
	"log"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/py-radicz/aoc25/utils"
)

const day int = 6

//var in string = `123 328  51 64
// 45 64  387 23
//  6 98  215 314
//*   +   *   +  `

type Formula struct {
	data []string
	op   string
}

func (f *Formula) PopOperator() {
	lastIdx := len(f.data) - 1
	f.op = f.data[lastIdx]
	f.data = f.data[:lastIdx]
}

func (f *Formula) CalcP1() (res int) {
	for _, num := range f.data {
		n, err := strconv.Atoi(strings.Trim(num, "X"))
		if err != nil {
			log.Fatal("failed to calculate formula")
		}

		switch f.op {
		case "*":
			if res == 0 {
				res = 1
			}
			res *= n
		case "+":
			res += n
		}

	}

	return res
}

func (f *Formula) CalcP2() (res int) {
	cols := len(f.data[0])
	nums := make(map[int]string, cols)

	for _, num := range f.data {
		for j := range cols {
			nums[j] += string(num[j])
		}
	}
	f.data = slices.Collect(maps.Values(nums))

	return f.CalcP1()
}

func FindColPos(rows []string) (res []int) {
	// column boundary is a column that have space in all rows in same place
	indexes := make(map[int]int, len(rows[0]))

	for _, row := range rows {
		for j, col := range row {
			if col == ' ' {
				indexes[j]++
			}
		}
	}

	for k, v := range indexes {
		if v == len(rows) {
			res = append(res, k)
		}
	}

	slices.Sort(res)
	return res
}

func ParseFormulas(rows []string, cols []int) []Formula {
	formulas := make([]Formula, len(cols)+1)

	for _, row := range rows {
		start := 0
		for j, idx := range cols {
			formulas[j].data = append(formulas[j].data, row[start:idx])
			start = idx + 1
			if j == len(cols)-1 {
				formulas[j+1].data = append(formulas[j+1].data, row[start:])
			}
		}

	}

	for i, f := range formulas {
		for j, num := range f.data {
			if strings.Contains(num, "*") || strings.Contains(num, "+") {
				formulas[i].data[j] = strings.Trim(num, " ")
				continue
			}
			formulas[i].data[j] = strings.ReplaceAll(num, " ", "X")
		}
	}
	return formulas
}

func DaySix() (partOne, partTwo int) {
	in, err := utils.GetInput(day)
	if err != nil {
		log.Fatalf("failed to load input for day %d", day)
	}
	rows := strings.Split(string(in), "\n")
	formulas := ParseFormulas(rows, FindColPos(rows))
	for _, f := range formulas {
		f.PopOperator()
		partOne += f.CalcP1()
		partTwo += f.CalcP2()
	}

	return
}
