package cryptosquare

import (
	"bytes"
	"math"
	"unicode"
)

const testVersion = 2

type Rect [][]string

func Encode(clearText string) string {
	text := normalize(clearText)
	if len(text) == 0 {
		return ""
	}
	rect := *buildRect(text)
	var b bytes.Buffer
	for col := range rect[0] {
		for row := range rect {
			b.WriteString(rect[row][col])
		}
		b.WriteString(" ")
	}
	return string(bytes.TrimRight(b.Bytes(), " "))
}

func normalize(input string) string {
	var b bytes.Buffer
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func buildRect(input string) *Rect {
	c, r := calcDimensions(input)
	rect := make(Rect, r)
	var inputIndex int
	for row := range rect {
		rect[row] = make([]string, c)
		for col := range rect[row] {
			if inputIndex >= len(input) {
				return &rect
			}
			rect[row][col] = string(input[inputIndex])
			inputIndex++
		}
	}

	return &rect
}

func calcDimensions(input string) (int, int) {
	guess := int(math.Floor(math.Sqrt(float64(len(input)))))
	c, r := guess, guess
	for {
		if c*r >= len(input) {
			return c, r
		}
		c += 1
		if c*r >= len(input) {
			return c, r
		}
		r += 1
	}
}
