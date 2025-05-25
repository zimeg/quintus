package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// Index acts as a landing page for the curious navigator
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%d-04-04T03:55:05Z", now.Moment(time.Now().UTC()).Year())
		return
	}
	tpl := `
<!DOCTYPE html>
<html>
    <head>
        <title>Quintus calendar</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <script src="https://cdn.tailwindcss.com"></script>
        <style type="text/tailwindcss">
        @layer base {
            a {
                @apply underline;
            }
            body > * {
                @apply font-mono max-w-lg mx-auto;
            }
            h1 {
                @apply pt-3 text-2xl;
            }
            nav {
                @apply pb-2 text-sm;
            }
        }
        </style>
    </head>
    <body>
        <h1>Quintus calendar</h1>
        <nav>
            <a
                href="https://o526.net/blog/post/five-day-week"
                target="_blank"
                title="the five day week"
            >post</a>
            <a
                href="https://github.com/zimeg/quintus"
                target="_blank"
                title="github repo"
            >code</a>
        </nav>
        <p>
            <span title="Quintus Time Server">QTS</span>
            ......
            <time id="quintus" hx-get="/now" hx-trigger="every 1s">
                {{.Time}}
            </time>
        </p>
        <p>
            <span title="Universal Coordinated Time">UTC</span>
            ......
            <time id="utc" hx-get="/utc" hx-trigger="every 1s">
                {{.UTC}}
            </time>
        </p>
        <script src="https://unpkg.com/htmx.org@2.0.2"></script>
    </body>
</html>`
	moment := now.Moment(time.Now().UTC()).ToString()
	t, err := template.New("index").Parse(tpl)
	if err != nil {
		fmt.Fprintf(w, "%s", moment)
		return
	}
	data := struct {
		UTC  string
		Time string
	}{
		UTC:  utc.ToString(),
		Time: moment,
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
}
