package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("Lengths must be the same")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
