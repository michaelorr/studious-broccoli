package foodchain

import "fmt"

const testVersion = 2

var creatures = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var reasons = []string{"", "It wriggled and jiggled and tickled inside her.\n",
	"How absurd to swallow a bird!\n",
	"Imagine that, to swallow a cat!\n",
	"What a hog, to swallow a dog!\n",
	"Just opened her throat and swallowed a goat!\n",
	"I don't know how she swallowed a cow!\n", "",
}

func Verse(i int) string {
	i = i - 1
	return fmt.Sprintf("%s%s%s%s",
		firstLine(i),
		reasons[i],
		progression(i),
		finalLine(i),
	)
}

func Verses(verses ...int) string {
	ret := ""
	for _, i := range verses {
		ret += Verse(i)
		if i != len(verses) {
			ret += "\n\n"
		}
	}
	return ret
}

func Song() string {
	numVerses := []int{1, 2, 3, 4, 5, 6, 7, 8}
	return Verses(numVerses...)
}

func firstLine(i int) string {
	return fmt.Sprintf("I know an old lady who swallowed a %s.\n", creatures[i])
}

func progression(i int) string {
	ret, clause := "", ""

	if i == 7 {
		return ret
	}
	for ; i > 0; i-- {
		if i == 2 {
			clause = " that wriggled and jiggled and tickled inside her"
		} else {
			clause = ""
		}
		ret += fmt.Sprintf("She swallowed the %s to catch the %s%s.\n", creatures[i], creatures[i-1], clause)
	}
	return ret
}

func finalLine(i int) string {
	if i == 7 {
		return "She's dead, of course!"
	} else {
		return "I don't know why she swallowed the fly. Perhaps she'll die."
	}
}
