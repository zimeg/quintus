package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

// Source links to the open source repository
func Source(w http.ResponseWriter, r *http.Request) {
	content := `<ul>
<li>Developed and pressed in San Francisco.</li>
<li>Built as <a href="https://github.com/zimeg/quintus" target="_blank" title="ssh git.o526.net -t quintus">open source</a> public schematics.</li>
<li>Based on prior <a href="https://o526.net/blog/post/five-day-week" target="_blank" title="the five day week">writings</a> to these topic.</li>
</ul>`
	if r.Header.Get("HX-Request") == "" {
		index(w, r, template.HTML(content))
		return
	}
	fmt.Fprint(w, content)
}
