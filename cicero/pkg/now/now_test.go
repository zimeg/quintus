package now

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMoment(t *testing.T) {
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
		"append the trailing nil period to the beginning of the year": {
			now:      time.Date(1999, 12, 27, 0, 0, 0, 0, time.UTC),
			expected: "2000-00-01T00:00:00Z",
			epoch:    time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"wait to start the next year with a nil period for resetting": {
			now:      time.Date(1999, 12, 31, 23, 59, 59, 99, time.UTC),
			expected: "2000-00-05T23:59:59Z",
			epoch:    time.Date(1999, 12, 31, 23, 59, 59, 99, time.UTC).Unix(),
		},
		"capture additions of a leap year increasing counting offsets": {
			now:      time.Date(2000, 12, 31, 23, 59, 59, 99, time.UTC),
			expected: "2001-00-06T23:59:59Z",
			epoch:    time.Date(2000, 12, 31, 23, 59, 59, 99, time.UTC).Unix(),
		},
		"confirm that leap years are counted using the correct offset": {
			now:      time.Date(2100, 12, 31, 23, 59, 59, 99, time.UTC),
			expected: "2101-00-05T23:59:59Z",
			epoch:    time.Date(2100, 12, 31, 23, 59, 59, 99, time.UTC).Unix(),
		},
		"find impossible dates repeat the date before if this happens": {
			now:      time.Date(2001, 2, 28, 0, 0, 0, 0, time.UTC),
			expected: "2001-02-29T00:00:00Z",
			epoch:    time.Date(2001, 2, 28, 0, 0, 0, 0, time.UTC).Unix(),
		},
		"check a known date around the time of development to be sure": {
			now:      time.Date(2024, 9, 21, 10, 26, 8, 5, time.UTC),
			expected: "2024-09-25T10:26:08Z",
			epoch:    time.Date(2024, 9, 25, 10, 26, 8, 5, time.UTC).Unix(),
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
