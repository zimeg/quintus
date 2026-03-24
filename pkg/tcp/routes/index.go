// Package routes has controllers for paths matching website pages with routers
// being defined elsewhere.
//
// https://quintus.sh
package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/zimeg/quintus/pkg/now"
	"github.com/zimeg/quintus/pkg/utc"
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
		if _, err := strconv.Atoi(r.URL.Path[1:]); err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%d-04-04T03:55:05Z", current.Year())
			return
		}
	}
	index(w, r, "")
}

// index renders the full page with an optional info section pre-filled
func index(w http.ResponseWriter, r *http.Request, info template.HTML) {
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
	year := current.Year()
	jump := fmt.Sprintf("%02d", current.Month())
	if y, err := strconv.Atoi(r.URL.Path[1:]); err == nil {
		year = y
		jump = "00"
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
		<script async src="https://plausible.io/js/pa-V5C8MVM8CutMwBhg8aipf.js"></script>
		<script>
			window.plausible=window.plausible||function(){(plausible.q=plausible.q||[]).push(arguments)},plausible.init=plausible.init||function(i){plausible.o=i||{}};
			plausible.init()
		</script>
	</head>
	<body>
		<form autocomplete="off">
			<header>
				<h1><a href="/">Quintus Calendars</a></h1>
				<nav>
					<a
						hx-get="/about"
						hx-target="#info"
						hx-swap="innerHTML"
					>about</a>
					<a
						hx-get="/shop"
						hx-target="#info"
						hx-swap="innerHTML"
					>shop</a>
					<a
						hx-get="/source"
						hx-target="#info"
						hx-swap="innerHTML"
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
				<section id="info">{{.Info}}</section>
			</header>
			<main>
				<table>
					<tbody>
						<tr>
							<th colspan="6">
								<a href="/{{ add .Year -1 }}">{{ add .Year -1 }}-00</a>
							</th>
						</tr>
						{{ .Curr }}
						{{ .Next }}
						<tr id="after"
							hx-get="/cal/{{ add .Year 2 }}"
							hx-trigger="revealed once"
							hx-swap="outerHTML"
						>
							<th colspan="6">{{ add .Year 2 }}-00</th>
						</tr>
					</tbody>
				</table>
			</main>
		</form>
		<script src="https://unpkg.com/htmx.org@2.0.2"></script>
		<script>
		if (!location.hash) {
			document.getElementById("{{ .Jump }}").scrollIntoView();
		}
		</script>
	</body>
</html>`
	t, err := template.New("index").Funcs(funcs).Parse(tpl)
	if err != nil {
		fmt.Fprintf(w, "%s", current.ToString())
		return
	}
	curr, err := calendar(year, current)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	next, err := calendar(year+1, current)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	data := struct {
		UTC  string
		Time now.Now
		Info template.HTML
		Year int
		Jump string
		Curr template.HTML
		Next template.HTML
	}{
		UTC:  utc.Current().In(location).ToString(),
		Time: current,
		Info: info,
		Year: year,
		Jump: jump,
		Curr: template.HTML(curr.String()),
		Next: template.HTML(next.String()),
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
}
