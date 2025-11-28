package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// Now returns the correct and current time according to the quintus calendar
func Now(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("timezone")
	timezone := "Etc/UTC"
	if cookie != nil {
		timezone = cookie.Value
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		location = time.UTC
	}
	fmt.Fprintf(w, "%s", now.Moment(utc.Current().In(location).ToTime()).ToString())
}
