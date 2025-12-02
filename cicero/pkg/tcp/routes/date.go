package routes

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// Date returns timers for the provided Gregorian date
func Date(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("timezone")
	timezone := "Etc/UTC"
	if cookie != nil {
		timezone = cookie.Value
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		location = time.UTC
	}
	current := now.Moment(utc.Current().In(location).ToTime())
	query := r.PathValue("date")
	gregorian, err := time.ParseInLocation("2006-01-02", query, location)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
		return
	}
	quintus := now.Moment(gregorian)
	if r.Header.Get("HX-Request") == "" {
		date := strings.Split(quintus.ToString(), "T")[0]
		fmt.Fprintf(w, "%s", date)
		return
	}
	if quintus.Year() == current.Year() &&
		quintus.Month() == current.Month() &&
		quintus.Date() == current.Date() {
		fmt.Fprintf(w, `
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
						%s
					</time>
				</p>
				<p>
					<span title="Universal Coordinated Time">UTC</span>
					......
					<time
						id="utc"
						hx-get="/utc"
						hx-params="none"
						hx-trigger="every 1s"
						hx-swap="innerHTML"
					>
						%s
					</time>
				</p>
			</article>`, current.ToString(), utc.Current().In(location).ToString())
		return
	}
	fmt.Fprintf(w, `
		<article 
			id="timers"
			hx-post="/timezone"
			hx-swap="outerHTML"
			hx-vals="js:{timezone: Intl.DateTimeFormat().resolvedOptions().timeZone}"
		>
			<p>
				<span title="Quintus Time Server">QTS</span>
				......
				<time id="quintus">
					%s
				</time>
			</p>
			<p>
				<span title="Universal Coordinated Time">UTC</span>
				......
				<time id="utc">
					%s
				</time>
			</p>
		</article>`, quintus.ToString(), utc.Moment(gregorian).ToString())
}
