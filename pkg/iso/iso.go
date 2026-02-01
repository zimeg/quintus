package iso

import (
	"time"
)

// Leaps decides if a year has a leap day
func Leaps(year int) bool {
	return 365 < time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
}
