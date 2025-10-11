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
	funcs := template.FuncMap{
		"add": func(i, j int) int {
			return i + j
		},
	}
	tpl := `
<!DOCTYPE html>
<html>
    <head>
        <title>Quintus calendar</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link href="./css/output.css" rel="stylesheet">
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
			<article id="timers">
				<p>
					<span title="Quintus Time Server">QTS</span>
					......
					<time id="quintus" hx-get="/now" hx-trigger="every 1s">
						{{.Time.ToString}}
					</time>
				</p>
				<p>
					<span title="Universal Coordinated Time">UTC</span>
					......
					<time id="utc" hx-get="/utc" hx-trigger="every 1s">
						{{.UTC}}
					</time>
				</p>
			</article>
		</header>
		<main>
			<form>
				<table>
					<tbody>
						<tr id="before"></tr>
						{{ .Prev }}
						<tr hx-get="/cal/{{ add .Time.Year -2 }}"
							hx-target="#before"
							hx-trigger="revealed once"
							hx-swap="outerHTML show:[id='{{ .Time.Year }}-00']:top"
						>
						</tr>
						{{ .Curr }}
						<tr hx-get="/cal/{{ add .Time.Year 2 }}"
							hx-target="#after"
							hx-trigger="revealed once"
							hx-swap="outerHTML"
						>
						</tr>
						{{ .Next }}
						<tr id="after"></tr>
					</tbody>
		  		</table>
			</form>
		</main>
		<script src="https://unpkg.com/htmx.org@2.0.2"></script>
		<script>
		htmx.on('htmx:afterSwap', (e) => {
			setTimeout(() => {
				document.querySelectorAll("tr[hx-disable]").forEach((el) => el.removeAttribute("hx-disable"));
				htmx.process(document.body);
			}, 120);
		});
		if (!location.hash) {
			document.getElementById("{{ .Time.Year }}-{{ .Time.Month }}").scrollIntoView();
		}
		</script>
	</body>
</html>`
	t, err := template.New("index").Funcs(funcs).Parse(tpl)
	if err != nil {
		fmt.Fprintf(w, "%s", current.ToString())
		return
	}
	prev, err := calendar(current.Year() - 1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	curr, err := calendar(current.Year())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	next, err := calendar(current.Year() + 1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	data := struct {
		UTC  string
		Time now.Now
		Prev template.HTML
		Curr template.HTML
		Next template.HTML
	}{
		UTC:  utc.ToString(),
		Time: current,
		Prev: template.HTML(prev.String()),
		Curr: template.HTML(curr.String()),
		Next: template.HTML(next.String()),
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
}
