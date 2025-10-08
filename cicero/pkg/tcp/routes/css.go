package routes

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed css/output.css
var css string

// CSS returns the compiled tailwind stylesheet
func CSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	fmt.Fprintf(w, "%s", css)
}
