package gigasecond

import (
	"math"
	"time"
)

const testVersion = 4

func AddGigasecond(bday time.Time) time.Time {
	return bday.Add(time.Duration(math.Pow(10, 9)) * time.Second)
}
