package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAbout confirms the about paragraph describes the quintus calendar
func TestAbout(t *testing.T) {
	tests := map[string]struct {
		htmx     bool
		contains string
	}{
		"describes the five day week of the quintus calendar system": {
			htmx:     true,
			contains: "five day week",
		},
		"mentions the thirty day months in the quintus calendar year": {
			htmx:     true,
			contains: "thirty day",
		},
		"mentions the liminal phase for the remaining calendar days": {
			htmx:     true,
			contains: "liminal phase",
		},
		"returns a full page with doctype for direct visits": {
			htmx:     false,
			contains: "<!DOCTYPE html>",
		},
		"includes the about content in the full page for direct visits": {
			htmx:     false,
			contains: "five day week",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/about", nil)
			require.NoError(t, err)
			if tt.htmx {
				req.Header.Set("HX-Request", "true")
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(About)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Contains(t, rr.Body.String(), tt.contains)
		})
	}
}
