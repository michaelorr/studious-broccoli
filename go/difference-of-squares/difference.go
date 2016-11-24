package diffsquares

import "math"

func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += math.Pow(float64(i), 2)
	}
	return int(sum)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
