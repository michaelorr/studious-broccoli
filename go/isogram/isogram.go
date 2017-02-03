package isogram

import "unicode"

const testVersion = 1

func IsIsogram(word string) bool {
	var mapping = make(map[rune]bool)
	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}
		str := unicode.ToLower(char)
		if dupe := mapping[str]; dupe {
			return false
		}
		mapping[str] = true
	}
	return true
}
