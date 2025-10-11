package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// Date returns timers for the provided Gregorian date
func Date(w http.ResponseWriter, r *http.Request) {
	current := now.Moment(utc.Now())
	query := r.PathValue("date")
	gregorian, err := time.Parse("2006-01-02", query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
		return
	}
	quintus := now.Moment(gregorian)
	if quintus.Year() == current.Year() &&
		quintus.Month() == current.Month() &&
		quintus.Date() == current.Date() {
		fmt.Fprintf(w, `<article id="timers">
			<p>
				<span title="Quintus Time Server">QTS</span>
					......
					<time id="quintus" hx-get="/now" hx-trigger="every 1s">
						%s
					</time>
				</p>
				<p>
					<span title="Universal Coordinated Time">UTC</span>
					......
					<time id="utc" hx-get="/utc" hx-trigger="every 1s">
						%s
					</time>
				</p>
			</article>`, current.ToString(), utc.ToString())
		return
	}
	fmt.Fprintf(w, `<article id="timers">
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
		</article>`, quintus.ToString(), gregorian.Format(time.RFC3339))
}
