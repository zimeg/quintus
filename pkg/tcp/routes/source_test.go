package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSource confirms the source paragraph links to the repository
func TestSource(t *testing.T) {
	tests := map[string]struct {
		htmx     bool
		contains string
	}{
		"mentions the development location of the quintus calendar": {
			htmx:     true,
			contains: "San Francisco",
		},
		"links to the github repository where the code is maintained": {
			htmx:     true,
			contains: "github.com/zimeg/quintus",
		},
		"returns a full page with doctype for direct visits": {
			htmx:     false,
			contains: "<!DOCTYPE html>",
		},
		"includes the source content in the full page for direct visits": {
			htmx:     false,
			contains: "San Francisco",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/source", nil)
			require.NoError(t, err)
			if tt.htmx {
				req.Header.Set("HX-Request", "true")
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Source)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Contains(t, rr.Body.String(), tt.contains)
		})
	}
}
