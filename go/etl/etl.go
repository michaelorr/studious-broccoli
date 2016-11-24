package etl

import "strings"

func Transform(m map[int][]string) map[string]int {
	res := make(map[string]int)
	for score, letters := range m {
		for _, letter := range letters {
			res[strings.ToLower(letter)] = score
		}
	}
	return res
}
