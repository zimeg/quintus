package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

// About describes the Quintus calendar
func About(w http.ResponseWriter, r *http.Request) {
	content := `<ul>
<li>A standard calendar with five day week.</li>
<li>Each month a six week total thirty day.</li>
<li>The remaining days makes liminal phase.</li>
</ul>`
	if r.Header.Get("HX-Request") == "" {
		index(w, r, template.HTML(content))
		return
	}
	fmt.Fprint(w, content)
}
