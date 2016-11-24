package raindrops

import (
	"bytes"
	"fmt"
)

const testVersion = 2

func Convert(n int) string {

	var buf bytes.Buffer

	if n%3 == 0 {
		buf.WriteString("Pling")
	}
	if n%5 == 0 {
		buf.WriteString("Plang")
	}
	if n%7 == 0 {
		buf.WriteString("Plong")
	}
	res := buf.String()
	if res == "" {
		res = fmt.Sprintf("%d", n)
	}
	return res
}
