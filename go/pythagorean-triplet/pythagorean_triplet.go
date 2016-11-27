package pythagorean

import (
	"math"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	var result []Triplet

	for a := min; a <= max; a++ {
		a_sq := math.Pow(float64(a), 2)
		for b := a + 1; b <= max; b++ {
			b_sq := math.Pow(float64(b), 2)
			for c := b + 1; c <= max; c++ {
				c_sq := math.Pow(float64(c), 2)
				if a_sq+b_sq == c_sq {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}
	return result
}

func Sum(p int) []Triplet {
	var result []Triplet
	for a := 1; a < p/3; a++ {
		a_sq := math.Pow(float64(a), 2)
		for b := a + 1; b < p; b++ {
			b_sq := math.Pow(float64(b), 2)
			c := p - (a + b)
			c_sq := math.Pow(float64(c), 2)
			if a_sq+b_sq == c_sq {
				result = append(result, Triplet{a, b, c})
			}
		}
	}
	return result
}
