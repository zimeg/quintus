package now

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoment(t *testing.T) {
	tz, err := time.LoadLocation("America/Phoenix")
	require.NoError(t, err)
	tests := map[string]struct {
		now      time.Time
		expected string
		epoch    int64
	}{
		"times match at the known epoch of the powerful unix machine": {
			now:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "1970-01-01T00:00:00Z",
			epoch:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"ensure the second month of quintus begins after thirty days": {
			now:      time.Date(1970, 1, 31, 0, 0, 0, 0, time.UTC),
			expected: "1970-02-01T00:00:00Z",
			epoch:    time.Date(1970, 2, 1, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"preserve the end of some current year before next leap year": {
			now:      time.Date(1999, 12, 26, 0, 0, 0, 0, time.UTC),
			expected: "1999-12-30T00:00:00Z",
			epoch:    time.Date(1999, 12, 30, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"append a trailing nil period to the beginning of next year": {
			now:      time.Date(1999, 12, 27, 0, 0, 0, 0, time.UTC),
			expected: "2000-00-01T00:00:00Z",
			epoch:    time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"increment the nil period according to sensible step counts": {
			now:      time.Date(1999, 12, 28, 0, 0, 0, 0, time.UTC),
			expected: "2000-00-02T00:00:00Z",
			epoch:    time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"match the expected nil period of with normal amount prior": {
			now:      time.Date(1999, 12, 31, 23, 59, 59, 99, time.UTC),
			expected: "2000-00-05T23:59:59Z",
			epoch:    time.Date(1999, 12, 31, 23, 59, 59, 99, time.UTC).Unix(),
		},
		"continue some year with a leaping nil period before start": {
			now:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "2000-00-06T00:00:00Z",
			epoch:    time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"begin the starts that leap in a mismatching step fashion": {
			now:      time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: "2000-01-01T00:00:00Z",
			epoch:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"find tomorrow is a single more than whatever started before": {
			now:      time.Date(2000, 1, 3, 0, 0, 0, 0, time.UTC),
			expected: "2000-01-02T00:00:00Z",
			epoch:    time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"explore an extra sunrise of the second month keeps aligned": {
			now:      time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: "2000-02-29T00:00:00Z",
			epoch:    time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"watch that extra sunrise has a noon without strange offset": {
			now:      time.Date(2000, 2, 29, 12, 34, 56, 789, time.UTC),
			expected: "2000-02-29T12:34:56Z",
			epoch:    time.Date(2000, 2, 29, 12, 34, 56, 789, time.UTC).Unix(),
		},
		"ideas of march repeat what a groundhog sees in cast shadow": {
			now:      time.Date(2000, 3, 1, 0, 0, 0, 0, time.UTC),
			expected: "2000-02-30T00:00:00Z",
			epoch:    time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"end a leap year on the expected end date with offsettings": {
			now:      time.Date(2000, 12, 26, 0, 0, 0, 0, time.UTC),
			expected: "2000-12-30T00:00:00Z",
			epoch:    time.Date(2000, 12, 30, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"start a next year the correct amount of offset expected days": {
			now:      time.Date(2000, 12, 27, 0, 0, 0, 0, time.UTC),
			expected: "2001-00-01T00:00:00Z",
			epoch:    time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"confirm leap days are kept with expected leap year offsets": {
			now:      time.Date(2000, 12, 31, 23, 59, 59, 99, time.UTC),
			expected: "2001-00-05T23:59:59Z",
			epoch:    time.Date(2000, 12, 31, 23, 59, 59, 99, time.UTC).Unix(),
		},
		"starting a standard count after leaps begins with schedule": {
			now:      time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "2001-01-01T00:00:00Z",
			epoch:    time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"find impossible dates repeat the date before if this happens": {
			now:      time.Date(2001, 2, 28, 0, 0, 0, 0, time.UTC),
			expected: "2001-02-29T00:00:00Z",
			epoch:    time.Date(2001, 2, 28, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"keep counting up on an irregular shortened month a marches": {
			now:      time.Date(2001, 3, 1, 0, 0, 0, 0, time.UTC),
			expected: "2001-02-30T00:00:00Z",
			epoch:    time.Date(2001, 2, 28, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"check a known date around the time of development to be sure": {
			now:      time.Date(2024, 9, 21, 10, 26, 8, 5, time.UTC),
			expected: "2024-09-24T10:26:08Z",
			epoch:    time.Date(2024, 9, 24, 10, 26, 8, 5, time.UTC).Unix(),
		},
		"confirm that leap years are counted using the correct offset": {
			now:      time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "2100-01-01T00:00:00Z",
			epoch:    time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"adjust output times to account for some timezone preference": {
			now:      time.Date(2026, 1, 1, 0, 0, 0, 0, tz),
			expected: "2026-01-01T00:00:00-07:00",
			epoch:    time.Date(2026, 1, 1, 0, 0, 0, 0, tz).Unix(),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Moment(tt.now)
			assert.Equal(t, tt.expected, actual.ToString())
			assert.Equal(t, tt.epoch, int64(actual.Epoch()))
		})
	}
}
