package pangram

const testVersion = 1

func IsPangram(text string) bool {
	found := make(map[rune]bool)
	for _, ch := range text {
		if (ch >= 65) && (ch <= 90) {
			ch = ch + 32
		}
		if (ch >= 97) && (ch <= 122) {
			found[ch] = true
		}
	}
	return len(found) == 26
}
