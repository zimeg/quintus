package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestShop confirms the shop paragraph offers printed calendars
func TestShop(t *testing.T) {
	tests := map[string]struct {
		htmx     bool
		contains string
	}{
		"shows the shop image for the stamped calendar print": {
			htmx:     true,
			contains: "shop.png",
		},
		"offers a pressing for ordering the current year": {
			htmx:     true,
			contains: "pressing for this current year",
		},
		"links to the stripe checkout for purchasing a calendar print": {
			htmx:     true,
			contains: "buy.stripe.com",
		},
		"returns a full page with doctype for direct visits": {
			htmx:     false,
			contains: "<!DOCTYPE html>",
		},
		"includes the shop content in the full page for direct visits": {
			htmx:     false,
			contains: "pressing for this current year",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/shop", nil)
			require.NoError(t, err)
			if tt.htmx {
				req.Header.Set("HX-Request", "true")
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Shop)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Contains(t, rr.Body.String(), tt.contains)
		})
	}
}
