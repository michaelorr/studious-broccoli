package sieve

const testVersion = 1

type candidate struct {
	value int
	valid bool
}

func Sieve(num int) []int {
	field := generateField(num)

	candidates := make(chan (int))
	go findCandidates(field, candidates)

	for candidate := range candidates {
		removeMultiples(candidate, field)
	}

	return findRemaining(field)
}

func generateField(num int) []candidate {
	field := make([]candidate, num-1)
	for i := range field {
		field[i] = candidate{value: i + 2, valid: true}
	}
	return field
}

func findCandidates(field []candidate, candidates chan (int)) {
	defer close(candidates)
	for _, possible := range field {
		if possible.valid == true {
			candidates <- possible.value
		}
	}
}

func removeMultiples(candidate int, field []candidate) {
	for i := 2; ; i++ {
		invalid := candidate * i
		if invalid >= len(field)+2 {
			return
		}
		field[invalid-2].valid = false
	}
}

func findRemaining(field []candidate) []int {
	var res []int
	for _, candidate := range field {
		if candidate.valid {
			res = append(res, candidate.value)
		}
	}
	return res
}
