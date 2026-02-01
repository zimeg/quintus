package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDate confirms the correct Quintus date is found for Gregorians
func TestDate(t *testing.T) {
	tests := map[string]struct {
		hx       bool
		date     string
		status   int
		exact    string
		expected []string
	}{
		"requests sent without the source from separate origin": {
			date:   "2022-02-22",
			status: 200,
			exact:  "2022-02-23",
		},
		"requests from the quintus calendar website use markup": {
			hx:     true,
			date:   "2022-02-22",
			status: 200,
			expected: []string{
				`<time id="quintus">`,
				"2022-02-23",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("GET /date/{date}", Date)
			req := httptest.NewRequest("GET", "/date/"+tt.date, nil)
			if tt.hx {
				req.Header.Set("HX-Request", "true")
			}
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			assert.Equal(t, tt.status, rr.Code)
			body := rr.Body.String()
			if tt.exact != "" {
				assert.Equal(t, tt.exact, body)
			}
			for _, a := range tt.expected {
				assert.Contains(t, body, a)
			}
		})
	}
}
