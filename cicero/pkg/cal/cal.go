package cal

import (
	"time"

	"github.com/zimeg/quintus/cicero/pkg/now"
)

// Calendar has months of a year
type Calendar struct {
	Year   int      // Year number
	Months [][]Date // Months have dates
}

// Date is the pairing of times
type Date struct {
	Gregorian time.Time
	Quintus   now.Now
}

// NewCalendar makes a relevant calendar
func NewCalendar(year int) Calendar {
	calendar := Calendar{
		Year:   year,
		Months: make([][]Date, 13),
	}
	g0 := time.Date(year-1, time.December, 27, 0, 0, 0, 0, time.UTC)
	g1 := time.Date(year-0, time.December, 26, 0, 0, 0, 0, time.UTC)
	for d := g0; !d.After(g1); d = d.AddDate(0, 0, 1) {
		q := now.Moment(d)
		m := q.Month()
		d := Date{
			Gregorian: d,
			Quintus:   q,
		}
		calendar.Months[m] = append(calendar.Months[m], d)
	}
	return calendar
}
