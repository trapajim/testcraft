package datagen

import (
	"time"
)

// Time returns a random time between 1970 and 2070.
func Time() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	sec := Rand().IntBetween(int(min), int(max))
	return time.Unix(int64(sec), 0)
}
