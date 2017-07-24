package anagram

import (
	"sort"
	"strings"
)

const testVersion = 2

func Detect(subject string, candidates []string) (result []string) {
	subject = strings.ToLower(subject)
	subjectSlice := strings.SplitN(subject, "", -1)
	sort.Strings(subjectSlice)
	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if candidateLower == subject {
			// identical strings are not anagrams
			continue
		}
		candidateSlice := strings.SplitN(candidateLower, "", -1)
		sort.Strings(candidateSlice)
		if strSliceEqual(candidateSlice, subjectSlice) {
			result = append(result, candidate)
		}
	}
	return result
}

func strSliceEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
