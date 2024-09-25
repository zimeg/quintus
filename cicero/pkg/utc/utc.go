package utc

import (
	"fmt"
	"time"
)

// Now returns the current time according the UTC
func Now() time.Time {
	return time.Now().UTC()
}

// ToString formats the current moment as a string
func ToString() string {
	moment := Now()
	return fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02dZ",
		moment.Year(),
		moment.Month(),
		moment.Day(),
		moment.Hour(),
		moment.Minute(),
		moment.Second(),
	)
}
