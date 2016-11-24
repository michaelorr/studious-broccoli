package letter

import (
	"sync"
)

func ConcurrentFrequency(words []string) FreqMap {
	results := make(chan FreqMap, len(words))

	var wg sync.WaitGroup
	for _, word := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			results <- Frequency(w)
		}(word)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalCount := FreqMap{}
	for freq := range results {
		for k, v := range freq {
			totalCount[k] += v
		}
	}

	return totalCount
}
