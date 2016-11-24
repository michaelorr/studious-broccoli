package house

import "fmt"

func Embed(relPhrase, nounPhrase string) string {
	return fmt.Sprintf("%s %s", relPhrase, nounPhrase)
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	return fmt.Sprintf("%s %s", subject, recurse(relPhrases, nounPhrase))
}

func recurse(relPhrases []string, nounPhrase string) string {
	if len(relPhrases) == 0 {
		return nounPhrase
	}
	return Embed(relPhrases[0], recurse(relPhrases[1:], nounPhrase))
}

func Song() string {
	subject := "This is"
	relPhrases := []string{
		"the horse and the hound and the horn\nthat belonged to",
		"the farmer sowing his corn\nthat kept",
		"the rooster that crowed in the morn\nthat woke",
		"the priest all shaven and shorn\nthat married",
		"the man all tattered and torn\nthat kissed",
		"the maiden all forlorn\nthat milked",
		"the cow with the crumpled horn\nthat tossed",
		"the dog\nthat worried",
		"the cat\nthat killed",
		"the rat\nthat ate",
		"the malt\nthat lay in",
	}
	nounPhrase := "the house that Jack built."

	res := subject + " " + nounPhrase

	for i := len(relPhrases) - 1; i >= 0; i-- {
		res += fmt.Sprintf("\n\n%s", Verse(subject, relPhrases[i:], nounPhrase))
	}
	return res
}
