package perfect

import "errors"

const testVersion = 1

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var (
	ErrOnlyPositive = errors.New("Positive number is required")
)

func Classify(n uint64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	sum := SumOfFactors(n)
	return compare(sum, n), nil
}

func compare(sum, n uint64) (class Classification) {
	if sum > n {
		class = ClassificationAbundant
	} else if sum == n {
		class = ClassificationPerfect
	} else {
		class = ClassificationDeficient
	}
	return
}

func SumOfFactors(n uint64) (sum uint64) {
	for i := uint64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return
}
