package grains

import (
	"errors"
	"math"
)

const testVersion = 1

func Total() (total uint64) {
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		total += s
	}
	return
}

func Square(i int) (uint64, error) {
	if (i <= 0) || (i > 64) {
		return 0, errors.New("invalid")
	}
	return uint64(math.Exp2(float64(i - 1))), nil
}
