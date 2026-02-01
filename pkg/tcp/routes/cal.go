package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zimeg/quintus/pkg/cal"
	"github.com/zimeg/quintus/pkg/now"
)

// calendar formats a table of dates to output
func calendar(year int, today now.Now) (w bytes.Buffer, err error) {
	funcs := template.FuncMap{
		"format": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
		"mod": func(i, j int) int {
			return i % j
		},
		"today": func(date now.Now) bool {
			return date.Year() == today.Year() &&
				date.Month() == today.Month() &&
				date.Date() == today.Date()
		},
	}
	tpl := `
{{- range $offset, $month := .Cal.Months -}}
{{- $d := index $month 0 -}}
{{- $id := printf "%d-%02d" $d.Quintus.Year $d.Quintus.Month }}
<tr>
	<th id="{{ $id }}" colspan="6">
		<a href="#{{ $id }}">{{ $id }}</a>
	</th>
</tr>
<tr>
	{{- range $num, $dates := $month }}
	<td
		title="{{ format $dates.Gregorian }}"
		id="{{ printf "Q%d-%02d-%02d" $dates.Quintus.Year $dates.Quintus.Month $dates.Quintus.Date }}"
	>
		<label>
			<p>
				<input
					type="radio"
					{{ if (today $dates.Quintus) }}checked{{ end }}
					name="date"
					value="{{ format $dates.Gregorian }}"
					hx-get="/date/{{ format $dates.Gregorian }}"
					hx-params="none"
					hx-target="#timers"
					hx-swap="outerHTML"
				>
				{{ if (lt $dates.Quintus.Date 10) }}&nbsp;{{ end }}{{ $dates.Quintus.Date }}
			</p>
		</label>
	</td>
	{{- if and (ne $dates.Quintus.Month 0) (eq (mod $dates.Quintus.Date 5) 0) (ne $dates.Quintus.Date 30) }}
</tr>
<tr>
	{{- end }}
	{{- end }}
</tr>
{{- end}}`
	t, err := template.New("cal").Funcs(funcs).Parse(tpl)
	if err != nil {
		return bytes.Buffer{}, err
	}
	calendar := cal.NewCalendar(year)
	data := struct {
		Year int
		Cal  cal.Calendar
	}{
		Year: year,
		Cal:  calendar,
	}
	err = t.Execute(&w, data)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return w, nil
}

// Cal responds to requests with a table
func Cal(w http.ResponseWriter, r *http.Request) {
	current := now.Moment(time.Now().UTC())
	year, err := func() (int, error) {
		switch r.PathValue("year") {
		case "":
			return current.Year(), nil
		default:
			return strconv.Atoi(r.PathValue("year"))
		}
	}()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%d-04-00T03:55:05Z", current.Year())
		return
	}
	cal, err := calendar(year, current)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%d-05-00T03:55:05Z", current.Year())
		return
	}
	switch {
	case year < current.Year():
		fmt.Fprintf(w, `
<tr id="before"></tr>
%s
<tr hx-get="/cal/%d"
	hx-disable
	hx-target="#before"
	hx-trigger="revealed once"
	hx-swap="outerHTML show:[id='%d-00']:top"
>
</tr>`,
			strings.TrimSpace(cal.String()),
			year-1,
			year+1,
		)
	case year > current.Year():
		fmt.Fprintf(w, `
<tr hx-get="/cal/%d"
	hx-disable
	hx-target="#after"
	hx-trigger="revealed once"
	hx-swap="outerHTML"
>
%s
<tr id="after"></tr>
		`,
			year+1,
			strings.TrimSpace(cal.String()),
		)
	default:
		fmt.Fprintf(w, `%s`, strings.TrimSpace(cal.String()))
	}
}
