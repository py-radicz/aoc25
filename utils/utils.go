/* Package utils */
package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	year = 2025
)

func GetInput(day int) ([][]byte, error) {
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
	return bytes.Fields(bodyText), nil
}

func Divmod(numerator, denominator int) (quotient, remainder int) {
	return numerator / denominator, numerator % denominator
}

func AtoiBytes(b []byte) int {
	var num int

	for _, c := range b {
		num = num*10 + int(c-'0')
	}
	return num
}
