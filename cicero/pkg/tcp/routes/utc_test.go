package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// TestRouteUTC confirms a current coordinated time is printed for requests
func TestRouteUTC(t *testing.T) {
	tests := map[string]struct {
		status   int
		timezone string
	}{
		"default to coordinated universal time without timezone": {
			status: 200,
		},
		"fallsback to the coordinated time for strange settings": {
			status:   200,
			timezone: "Moon/Luna",
		},
		"adjusts to show the location setting print as current": {
			status:   200,
			timezone: "America/Los_Angeles",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/utc", nil)
			require.NoError(t, err)
			if tt.timezone != "" {
				cookie := &http.Cookie{
					Name:  "timezone",
					Value: tt.timezone,
				}
				req.AddCookie(cookie)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(UTC)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, tt.status, rr.Code)
			location, err := time.LoadLocation(tt.timezone)
			if err != nil {
				location = time.UTC
			}
			expected := utc.Current().In(location).ToString()
			assert.Equal(t, rr.Body.String(), expected)
		})
	}
}
