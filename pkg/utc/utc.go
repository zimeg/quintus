// Package utc represents the "Coordinated Universal Time" for converting from
// another calendar implementation.
//
// https://en.wikipedia.org/wiki/Coordinated_Universal_Time
package utc

import (
	"fmt"
	"time"
)

// UTC has a moment in time
type UTC struct {
	moment time.Time
}

// Current returns a current time according the UTC
func Current() *UTC {
	return &UTC{
		moment: time.Now().UTC(),
	}
}

// Moment casts a time to some UTC representation
func Moment(t time.Time) *UTC {
	return &UTC{
		moment: t,
	}
}

// In sets the timezone to consider this cordination in
func (u *UTC) In(loc *time.Location) *UTC {
	u.moment = u.moment.In(loc)
	return u
}

// ToTime returns the time.
func (u UTC) ToTime() time.Time {
	return u.moment
}

// ToString formats the current moment as a string
func (u UTC) ToString() string {
	timezone := "Z"
	if u.moment.Location().String() != "UTC" &&
		u.moment.Location().String() != "Etc/UTC" {
		_, sec := time.Now().In(u.moment.Location()).Zone()
		sign := "+"
		if sec < 0 {
			sign = "-"
			sec = -sec
		}
		h := sec / 3600
		m := (sec % 3600) / 60
		timezone = fmt.Sprintf("%s%02d:%02d", sign, h, m)
	}
	return fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02d%s",
		u.moment.Year(),
		u.moment.Month(),
		u.moment.Day(),
		u.moment.Hour(),
		u.moment.Minute(),
		u.moment.Second(),
		timezone,
	)
}
