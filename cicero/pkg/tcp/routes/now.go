package routes

import (
	"fmt"
	"net/http"

	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// Now returns the correct and current time according to the quintus calendar
func Now(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", now.Moment(utc.Now()).ToString())
}
