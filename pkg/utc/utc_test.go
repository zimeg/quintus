package utc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrent(t *testing.T) {
	current := Current()
	assert.GreaterOrEqual(t, current.ToTime().Year(), 2025)
}

func TestMoment(t *testing.T) {
	tests := map[string]struct {
		timezone string
		date     string
		expected string
	}{
		"Etc/UTC": {
			timezone: "Etc/UTC",
			date:     "2025-11-05",
			expected: "2025-11-05T00:00:00Z",
		},
		"America/Phoenix": {
			timezone: "America/Phoenix",
			date:     "2025-11-05",
			expected: "2025-11-04T17:00:00-07:00",
		},
		"Australia/Darwin": {
			timezone: "Australia/Darwin",
			date:     "2025-11-05",
			expected: "2025-11-05T09:30:00+09:30",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			location, err := time.LoadLocation(tt.timezone)
			require.NoError(t, err)
			parsed, err := time.Parse("2006-01-02", tt.date)
			require.NoError(t, err)
			utc := Moment(parsed)
			assert.Equal(t, tt.expected, utc.In(location).ToString())
		})
	}
}
