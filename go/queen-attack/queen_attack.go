package queenattack

import (
	"errors"
	"strconv"
)

const testVersion = 1

type position struct {
	x int
	y int
}

func CanQueenAttack(w, b string) (bool, error) {
	white, err := parse(w)
	if err != nil {
		return false, err
	}
	black, err := parse(b)
	if err != nil {
		return false, err
	}

	if white == black {
		return false, errors.New("invalid")
	}

	return canAttack(white, black), nil
}

func canAttack(white, black position) bool {
	if black.x == white.x {
		return true
	}
	if black.y == white.y {
		return true
	}

	blackDiag := make([]position, 0)

	i := black.x
	j := black.y
	for (i >= 0) && (j >= 0) {
		i--
		j--
	}
	for (i < 8) && (j < 8) {
		blackDiag = append(blackDiag, position{i, j})
		i++
		j++
	}

	i = black.x
	j = black.y
	for (i >= 0) && (j >= 0) {
		i--
		j++
	}
	for (i < 8) && (j < 8) {
		blackDiag = append(blackDiag, position{i, j})
		i++
		j--
	}

	for _, v := range blackDiag {
		if v == white {
			return true
		}
	}

	return false
}

func parse(s string) (position, error) {
	if len(s) != 2 {
		return position{}, errors.New("invalid")
	}
	x := int(rune(s[0]) - 'a')
	y, _ := strconv.Atoi(s[1:])
	if x > 8 || y > 8 {
		return position{}, errors.New("invalid")
	}
	return position{x, y}, nil
}
