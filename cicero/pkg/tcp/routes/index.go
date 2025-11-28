// Package routes has controllers for paths matching website pages with routers
// being defined elsewhere.
//
// https://quintus.sh
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
	cookie, _ := r.Cookie("timezone")
	timezone := "Etc/UTC"
	if cookie != nil {
		timezone = cookie.Value
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		location = time.UTC
	}
	current := now.Moment(time.Now().In(location))
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
		<title>Quintus Calendars</title>
		<meta charset="UTF-8">
		<meta name="description" content="The Quintus calendar is an alternative to the irregular Gregorian calendar: consistent 5 day weeks make equal 30 day months." />
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="apple-touch-icon" href="/apple-touch-icon.png" sizes="180x180">
		<link rel="canonical" href="https://quintus.sh/" />
		<link rel="icon" href="/favicon.ico" sizes="any">
		<link rel="icon" href="/favicon-32x32.png" sizes="32x32">
		<link rel="manifest" href="/manifest.webmanifest">
		<link href="./css/output.css" rel="stylesheet">
		<meta property="og:title" content="Quintus Calendars" />
		<meta property="og:description" content="The Quintus calendar is an alternative to the irregular Gregorian calendar: consistent 5 day weeks make equal 30 day months." />
		<meta property="og:image" content="https://o526.net/blog/note/ff8bd197/calendar.png" />
		<meta property="og:image:alt" content="A stamped marking of upcoming month with leaf" />
		<meta property="og:type" content="website" />
		<meta property="og:url" content="https://quintus.sh/" />
	</head>
	<body>
		<form autocomplete="off">
			<header>
				<h1>Quintus Calendars</h1>
				<nav>
					<a
						href="https://o526.net/blog/post/five-day-week"
						target="_blank"
						title="the five day week"
					>about</a>
					<a
						href="https://buy.stripe.com/cNiaEZ0G56p5eQFgiB9EI00"
						target="_blank"
						title="checkout"
					>shop</a>
					<a
						href="https://github.com/zimeg/quintus"
						target="_blank"
						title="github repo"
					>source</a>
				</nav>
				<article
					id="timers"
					hx-post="/timezone"
					hx-swap="outerHTML"
					hx-vals="js:{timezone: Intl.DateTimeFormat().resolvedOptions().timeZone}"
				>
					<p>
						<span title="Quintus Time Server">QTS</span>
						......
						<time
							id="quintus"
							hx-get="/now"
							hx-params="none"
							hx-swap="innerHTML"
							hx-trigger="every 1s"
						>
							{{.Time.ToString}}
						</time>
					</p>
					<p>
						<span title="Universal Coordinated Time">UTC</span>
						......
						<time
							id="utc"
							hx-get="/utc"
							hx-params="none"
							hx-swap="innerHTML"
							hx-trigger="every 1s"
						>
							{{.UTC}}
						</time>
					</p>
				</article>
			</header>
			<main>
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
			</main>
		</form>
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
	prev, err := calendar(current.Year()-1, current)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	curr, err := calendar(current.Year(), current)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	next, err := calendar(current.Year()+1, current)
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
		UTC:  utc.Current().In(location).ToString(),
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
