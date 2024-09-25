package now

import (
	"fmt"
	"time"
)

// Now is this moment
type Now struct {
	year   int
	month  int
	date   int
	hour   int
	minute int
	second int
}

// Moment realizes the provided utc time in the standard quintus format
func Moment(utc time.Time) Now {
	return Now{
		year:   utc.Year(),
		month:  utc.YearDay() / 30,
		date:   utc.YearDay() % 30,
		hour:   utc.Hour(),
		minute: utc.Minute(),
		second: utc.Second(),
	}
}

// Epoch represents the number of seconds since the Unix epoch for the current
// moment in the Quintus format but offset to match the Gregorian calendar when
// converted to ISO 8601 standards
//
// Dates that the Gregorian calendar cannot show are waited for on the previous
// date until the next possible date arrives
func (n *Now) Epoch() uint64 {
	var month time.Month
	var date int
	if n.month != 12 {
		month = time.Month(n.month + 1)
		date = n.date
	} else {
		month = time.Month(12)
		date = 31
	}
	conversion := time.Date(
		n.year,
		month,
		date,
		n.hour,
		n.minute,
		n.second,
		0,
		time.UTC,
	)
	for conversion.Month() != month {
		conversion = conversion.Add(time.Duration(-24) * time.Hour)
	}
	return uint64(conversion.Unix())
}

// Offset adds the seconds since the NTP epoch 1900-01-01 00:00:00 to the valid
// Unix epoch of now
func (n *Now) Offset() uint64 {
	return uint64(2208988800+n.Epoch()) << 32
}

// ToString converts internal representations of now into written string
func (n Now) ToString() string {
	return fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02dZ",
		n.year,
		n.month+1,
		n.date,
		n.hour,
		n.minute,
		n.second,
	)
}
