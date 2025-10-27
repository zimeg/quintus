package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// Timezone swaps between server time and client time for following requests
func Timezone(w http.ResponseWriter, r *http.Request) {
	moment := now.Moment(time.Now().UTC())
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%d-04-05T03:55:05Z", moment.Year())
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	found, err := r.Cookie("timezone")
	location := r.FormValue("timezone")
	if err != nil {
		if err != http.ErrNoCookie {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
			return
		}
	} else {
		if found.Value != "Etc/UTC" {
			location = "Etc/UTC"
		}
	}
	existing, err := time.LoadLocation(found.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	requested, err := time.LoadLocation(location)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", moment.Year())
		return
	}
	current := now.Moment(time.Now().In(requested))
	cookie := &http.Cookie{
		Name:     "timezone",
		Value:    location,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	date := r.FormValue("date")
	before, err := time.Parse("2006-01-02", date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
		return
	}
	clock := time.Now().In(existing)
	previous, err := time.ParseInLocation("2006-01-02", date, existing)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
		return
	}
	previousa := time.Date(previous.Year(), previous.Month(), previous.Day(), clock.Hour(), clock.Minute(), clock.Second(), clock.Nanosecond(), existing)
	gregorian := time.Now().In(requested)

	templates := bytes.Buffer{}
	if previousa.UTC().Equal(clock.UTC()) {
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
			fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
			return
		}
		data := struct {
			NewGregorian time.Time
			NewQuintus   now.Now
			OldGregorian time.Time
			OldQuintus   now.Now
		}{
			NewGregorian: time.Now().In(requested),
			NewQuintus:   current,
			OldGregorian: before,
			OldQuintus:   now.Moment(before),
		}
		err = t.Execute(&templates, data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
			return
		}
	}

	if (before.Year() == gregorian.Year() &&
		before.Month() == gregorian.Month() &&
		before.Day() == gregorian.Day()) || previousa.UTC().Equal(clock.UTC()) {
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
			current.ToString(),
			utc.Current().In(requested).ToString(),
			templates.String(),
		)
		return
	}
	static, err := time.ParseInLocation("2006-01-02", date, requested)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
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
		now.Moment(static).ToString(),
		utc.Moment(static).ToString(),
		templates.String(),
	)
}
