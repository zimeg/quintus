package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zimeg/quintus/pkg/now"
)

func TestIndex(t *testing.T) {
	current := now.Moment(time.Now())
	tests := map[string]struct {
		expected []string
		status   int
	}{
		"headers": {
			expected: []string{
				"<title>Quintus Calendars</title>",
				"<h1>Quintus Calendars</h1>",
			},
			status: 200,
		},
		"scroll": {
			expected: []string{
				fmt.Sprintf(
					"document.getElementById(\"%d-%02d\").scrollIntoView();",
					current.Year(),
					current.Month(),
				),
			},
			status: 200,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("GET /", Index)
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)
			assert.Equal(t, tt.status, rr.Code)
			body := rr.Body.String()
			for _, a := range tt.expected {
				assert.Contains(t, body, a)
			}
		})
	}
}
