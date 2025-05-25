package now

import (
	"fmt"
	"time"

	"github.com/zimeg/quintus/cicero/pkg/iso"
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
	offset := utc.YearDay() - 1
	if iso.Leaps(utc.Year()) {
		offset -= 1
	}
	now := Now{
		year:   utc.Year(),
		month:  (offset / 30),
		date:   (offset % 30) + 1,
		hour:   utc.Hour(),
		minute: utc.Minute(),
		second: utc.Second(),
	}
	switch now.month {
	case 12:
		now.month = 0
		now.year += 1
	case 0:
		if iso.Leaps(now.year) && now.date == 0 {
			now.date = 6
			break
		}
		fallthrough
	default:
		now.month += 1
	}
	return now
}

// Epoch represents the number of seconds since the Unix epoch for the current
// moment in the Quintus format but offset to match the Gregorian calendar when
// converted to ISO 8601 standards
//
// Dates that the Gregorian calendar cannot show are waited for on the previous
// date until the next possible date arrives
func (n Now) Epoch() uint64 {
	year := n.year
	month := time.Month(n.month)
	date := n.date
	switch n.month {
	case 0:
		year = n.year - 1
		month = time.Month(12)
		date = 31
	}
	conversion := time.Date(
		year,
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
func (n Now) Offset() uint64 {
	return uint64(2208988800+n.Epoch()) << 32
}

// ToString converts internal representations of now into written string
func (n Now) ToString() string {
	return fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02dZ",
		n.year,
		n.month,
		n.date,
		n.hour,
		n.minute,
		n.second,
	)
}
