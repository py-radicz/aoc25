/* Package utils */
package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const year = 2025

type Puzzle struct {
	Func  func() (int, int)
	Day   string
	Part1 int
	Part2 int
}

func (p *Puzzle) Run() {
	start := time.Now()
	p1, p2 := p.Func()
	p.Part1, p.Part2 = p1, p2
	fullname := runtime.FuncForPC(reflect.ValueOf(p.Func).Pointer()).Name()
	parts := strings.Split(fullname, ".")
	p.Day = parts[len(parts)-1]
	fmt.Println("Total execution in: ", time.Since(start))
}

func (p *Puzzle) String() string {
	return fmt.Sprintf("\n%s solution\nPart1: %d\nPart2: %d", p.Day, p.Part1, p.Part2)
}

func NewPuzzle(f func() (int, int)) *Puzzle {
	return &Puzzle{Func: f, Part1: 0, Part2: 0}
}

func GetInput(day int) ([]byte, error) {
	start := time.Now()
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("cookie", os.Getenv("COOKIE"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Input loaded in: ", time.Since(start))
	return bytes.Trim(bodyText, " \n\r"), nil
}

func AtoiBytes(b []byte) int {
	var num int

	for _, c := range b {
		num = num*10 + int(c-'0')
	}
	return num
}

func KeyWithMinValue[K comparable, V ~int | ~int64 | ~float64](m map[K]V) (key K, value V, found bool) {
	for k, v := range m {
		if !found || v < value {
			key = k
			value = v
			found = true
		}
	}
	return
}
