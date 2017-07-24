package lsproduct

import "errors"
import "strconv"

const testVersion = 5

func LargestSeriesProduct(digits string, span int) (int, error) {
	if err := validateArgs(digits, span); err != nil {
		return 0, err
	}

	intSlice, err := sliceAndConvert(digits)
	if err != nil {
		return 0, err
	}

	products := localProducts(intSlice, span)
	return findMax(products), nil
}

func findMax(ints []int) int {
	var max int
	for _, possible := range ints {
		if possible > max {
			max = possible
		}
	}
	return max
}

func localProducts(ints []int, span int) []int {
	res := make([]int, len(ints)-span+1)
	for i := 0; i <= len(ints)-span; i++ {
		res[i] = multiplySegment(ints[i : i+span])
	}
	return res
}

func multiplySegment(ints []int) int {
	product := 1
	for _, num := range ints {
		product *= num
	}
	return product
}

func validateArgs(digits string, span int) (err error) {
	if span > len(digits) || span < 0 {
		err = errors.New("invalid span")
	}
	return
}

func sliceAndConvert(digits string) ([]int, error) {
	ints := make([]int, len(digits))
	for i, r := range digits {
		val, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, err
		}
		ints[i] = val
	}
	return ints, nil
}
