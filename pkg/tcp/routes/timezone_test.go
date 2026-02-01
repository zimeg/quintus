package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestTimezone converts between an origin and destination convention
func TestTimezone(t *testing.T) {
	tests := map[string]struct {
		cookie   string
		timezone string
		date     string
		status   int
		expected []string
	}{
		"without a cookie set arrive to california on past date": {
			timezone: "America/Los_Angeles",
			date:     "2022-02-22",
			status:   200,
			expected: []string{
				"2022-02-22T00:00:00-08:00",
			},
		},
		"with a cookie of location revert back to coordinations": {
			cookie:   "America/Los_Angeles",
			timezone: "America/Los_Angeles",
			date:     "2022-02-22",
			status:   200,
			expected: []string{
				"2022-02-22T00:00:00Z",
			},
		},
		"with a cookie of coordinations find again the location": {
			cookie:   "Etc/UTC",
			timezone: "America/Los_Angeles",
			date:     "2022-02-22",
			status:   200,
			expected: []string{
				"2022-02-22T00:00:00-08:00",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			form := url.Values{
				"timezone": {tt.timezone},
				"date":     {tt.date},
			}
			req, err := http.NewRequest("POST", "/timezone", strings.NewReader(form.Encode()))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			if tt.cookie != "" {
				cookie := &http.Cookie{
					Name:  "timezone",
					Value: tt.cookie,
				}
				req.AddCookie(cookie)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Timezone)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, tt.status, rr.Code)
			body := rr.Body.String()
			for _, a := range tt.expected {
				assert.Contains(t, body, a)
			}
		})
	}
}
