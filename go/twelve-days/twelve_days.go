package twelve

import "fmt"

const testVersion = 1

var phrases = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves, and ",
	3:  "three French Hens, ",
	4:  "four Calling Birds, ",
	5:  "five Gold Rings, ",
	6:  "six Geese-a-Laying, ",
	7:  "seven Swans-a-Swimming, ",
	8:  "eight Maids-a-Milking, ",
	9:  "nine Ladies Dancing, ",
	10: "ten Lords-a-Leaping, ",
	11: "eleven Pipers Piping, ",
	12: "twelve Drummers Drumming, ",
}

var verseNum = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var template = "On the %s day of Christmas my true love gave to me, %s."

func Song() string {
	response := ""
	for verse := 1; verse <= 12; verse++ {
		response = fmt.Sprintf("%s%s\n", response, Verse(verse))
	}
	return response
}

func Verse(verse int) string {
	return fmt.Sprintf(template, verseNum[verse], buildSuffix(verse))
}

func buildSuffix(verse int) string {
	if verse == 0 {
		return ""
	}
	return fmt.Sprintf("%s%s", phrases[verse], buildSuffix(verse-1))
}
