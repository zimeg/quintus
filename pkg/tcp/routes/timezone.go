package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/zimeg/quintus/pkg/now"
	"github.com/zimeg/quintus/pkg/utc"
)

// Timezone swaps between server time and client time for following requests
func Timezone(w http.ResponseWriter, r *http.Request) {
	moment := now.Moment(time.Now().UTC())
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%d-04-05T03:55:05Z", moment.Year())
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	saved, err := r.Cookie("timezone")
	if err != nil {
		if err != http.ErrNoCookie {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
			return
		}
	}
	from := "Etc/UTC"
	if saved != nil {
		from = saved.Value
	}
	origin, err := time.LoadLocation(from)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	to := r.FormValue("timezone")
	if from != "Etc/UTC" {
		to = "Etc/UTC"
	}
	destination, err := time.LoadLocation(to)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	cookie := &http.Cookie{
		Name:     "timezone",
		Value:    to,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	date := r.FormValue("date")
	dated, err := time.Parse("2006-01-02", date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	departure, err := time.ParseInLocation("2006-01-02", date, origin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	arrival, err := time.ParseInLocation("2006-01-02", date, destination)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	outbound := time.Now().In(origin)
	departing := time.Date(departure.Year(), departure.Month(), departure.Day(), outbound.Hour(), outbound.Minute(), outbound.Second(), outbound.Nanosecond(), origin).Equal(outbound)
	inbound := time.Now().In(destination)
	arriving := time.Date(arrival.Year(), arrival.Month(), arrival.Day(), inbound.Hour(), inbound.Minute(), inbound.Second(), inbound.Nanosecond(), destination).Equal(inbound)
	templates := bytes.Buffer{}
	if departing && !arriving {
		funcs := template.FuncMap{
			"format": func(t time.Time) string {
				return t.Format("2006-01-02")
			},
		}
		tpl := `
			<template>
				<td
					title="{{ format .NewGregorian }}"
					id="{{ printf "Q%d-%02d-%02d" .NewQuintus.Year .NewQuintus.Month .NewQuintus.Date }}"
					hx-swap-oob="true"
				>
					<label>
						<p>
							<input
								type="radio"
								checked
								name="date"
								value="{{ format .NewGregorian }}"
								hx-get="/date/{{ format .NewGregorian }}"
								hx-params="none"
								hx-target="#timers"
								hx-swap="outerHTML"
							>
							{{ if (lt .NewQuintus.Date 10) }}&nbsp;{{ end }}{{ .NewQuintus.Date }}
						</p>
					</label>
				</td>
				<td
					title="{{ format .OldGregorian }}"
					id="{{ printf "Q%d-%02d-%02d" .OldQuintus.Year .OldQuintus.Month .OldQuintus.Date }}"
					hx-swap-oob="true"
				>
					<label>
						<p>
							<input
								type="radio"
								name="date"
								value="{{ format .OldGregorian }}"
								hx-get="/date/{{ format .OldGregorian }}"
								hx-params="none"
								hx-target="#timers"
								hx-swap="outerHTML"
							>
							{{ if (lt .OldQuintus.Date 10) }}&nbsp;{{ end }}{{ .OldQuintus.Date }}
						</p>
					</label>
				</td>
			</template>`
		t, err := template.New("cal").Funcs(funcs).Parse(tpl)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
			return
		}
		data := struct {
			NewGregorian time.Time
			NewQuintus   now.Now
			OldGregorian time.Time
			OldQuintus   now.Now
		}{
			NewGregorian: time.Now().In(destination),
			NewQuintus:   now.Moment(time.Now().In(destination)),
			OldGregorian: dated,
			OldQuintus:   now.Moment(dated),
		}
		err = t.Execute(&templates, data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
			return
		}
	}
	if arriving || departing {
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
						hx-swap="innerHTML"
						hx-trigger="every 1s"
					>
						%s
					</time>
				</p>
			</article>
			%s`,
			now.Moment(time.Now().In(destination)).ToString(),
			utc.Current().In(destination).ToString(),
			templates.String(),
		)
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
		</article>
		%s`,
		now.Moment(arrival).ToString(),
		utc.Moment(arrival).ToString(),
		templates.String(),
	)
}
