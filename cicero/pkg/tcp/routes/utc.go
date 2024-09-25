package routes

import (
	"fmt"
	"net/http"

	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// UTC returns the universal time in the gregorian representation
func UTC(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", utc.ToString())
}
