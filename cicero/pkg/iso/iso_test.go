package iso

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaps(t *testing.T) {
	tests := map[string]struct {
		year     int
		expected bool
	}{
		"2000 leaps": {
			year:     2000,
			expected: true,
		},
		"2001 hops": {
			year:     2001,
			expected: false,
		},
		"2002 hops": {
			year:     2002,
			expected: false,
		},
		"2004 leaps": {
			year:     2004,
			expected: true,
		},
		"2100 hops": {
			year:     2100,
			expected: false,
		},
		"2200 hops": {
			year:     2200,
			expected: false,
		},
		"2400 leaps": {
			year:     2400,
			expected: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Leaps(tt.year)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
