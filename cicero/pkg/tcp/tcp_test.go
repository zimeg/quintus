package tcp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/tcp/routes"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// TestRouteIndex ensures important information is shown on an initial page load
// but leaves route testing to ci
func TestRouteIndex(t *testing.T) {
	tests := map[string]struct {
		path     string
		status   int
		expected []string
	}{
		"prints the most relevant information related to timing": {
			path:   "/",
			status: 200,
			expected: []string{
				now.Moment(utc.Now()).ToString(),
				utc.ToString(),
			},
		},
		"errors with an exit code when visiting an unknown path": {
			path:   "gregorian",
			status: 404,
			expected: []string{
				fmt.Sprintf("%d-04-04T03:55:05Z", time.Now().Year()),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(routes.Index)
			handler.ServeHTTP(rr, req)
			for _, str := range tt.expected {
				assert.Equal(t, tt.status, rr.Code)
				assert.Contains(t, rr.Body.String(), str)
			}
		})
	}
}
