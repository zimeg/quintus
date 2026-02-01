package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimeg/quintus/pkg/now"
	"github.com/zimeg/quintus/pkg/utc"
)

// TestRouteNow confirms a current quintus time is printed for requests
func TestRouteNow(t *testing.T) {
	tests := map[string]struct {
		status   int
		timezone string
	}{
		"default to a coordinated quintus time without timezone": {
			status: 200,
		},
		"fallsback to a coordinated quintus in strange settings": {
			status:   200,
			timezone: "Moon/Luna",
		},
		"adjusts to show the quintus location setting sent with": {
			status:   200,
			timezone: "America/Los_Angeles",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/now", nil)
			require.NoError(t, err)
			if tt.timezone != "" {
				cookie := &http.Cookie{
					Name:  "timezone",
					Value: tt.timezone,
				}
				req.AddCookie(cookie)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Now)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, tt.status, rr.Code)
			location, err := time.LoadLocation(tt.timezone)
			if err != nil {
				location = time.UTC
			}
			expected := now.Moment(utc.Current().In(location).ToTime()).ToString()
			assert.Equal(t, rr.Body.String(), expected)
		})
	}
}
