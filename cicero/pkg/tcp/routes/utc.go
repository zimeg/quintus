package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// UTC returns the universal time in the gregorian representation
func UTC(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("timezone")
	timezone := "Etc/UTC"
	if cookie != nil {
		timezone = cookie.Value
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		location = time.UTC
	}
	fmt.Fprintf(w, "%s", utc.Current().In(location).ToString())
}
