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
	current := now.Moment(time.Now().UTC())
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%d-04-04T03:55:05Z", current.Year())
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
                @apply font-mono;
				@apply max-w-lg;
				@apply mx-auto;
            }
            h1 {
                @apply pt-3;
				@apply text-2xl;
            }
			header {
				@apply bg-white;
				@apply border-b;
				@apply h-36;
				@apply left-0;
				@apply p-2;
				@apply sticky;
				@apply top-0;
				@apply w-full;
				@apply lg:fixed;
				@apply lg:h-full;
				@apply lg:w-96;
			}
			main {
				@apply lg:ml-96;
			}
			table {
				@apply table-fixed;
				@apply w-full;
			}
			td {
				@apply cursor-default;
				@apply text-right;
			}
			th {
				@apply cursor-pointer;
				@apply font-semibold;
				@apply scroll-mt-36;
				@apply text-right;
				@apply lg:scroll-mt-0;
			}
			time {
				@apply text-nowrap;
			}
            nav {
                @apply pb-2;
				@apply text-sm;
            }
        }
        </style>
    </head>
    <body>
		<header>
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
		</header>
		<main>
			<table>
				<tbody>{{.Cal}}</tbody>
		  	</table>
		</main>
        <script src="https://unpkg.com/htmx.org@2.0.2"></script>
    </body>
</html>`
	t, err := template.New("index").Parse(tpl)
	if err != nil {
		fmt.Fprintf(w, "%s", current.ToString())
		return
	}
	cal, err := calendar(current.Year())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	data := struct {
		UTC  string
		Time string
		Cal  template.HTML
	}{
		Cal:  template.HTML(cal.String()),
		UTC:  utc.ToString(),
		Time: current.ToString(),
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
}
