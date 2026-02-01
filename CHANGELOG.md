# Changelog

All changes that happen to this project will be documented in this file.

Updates follow a [conventional commit][commits] style and the project is
versioned with [calendar versioning][calver] using the [quintus][quintus]
calendar.

## Changes

- build: move hosting to self managed computers the ordering machine 2026-02-01
- fix: scroll to the more single digit months with left padded zeros 2026-01-28
- chore: include starting scripts for plausible page view measurment 2025-12-07
- build: recreate snapshot title after rebuilding the server package 2025-12-07
- chore: count page views for purposes of measuring quintus belivers 2025-12-07
- build: assign random but cute names to servers on the amazon farms 2025-12-06
- fix: permit access to a converted date from websites across origin 2025-12-06
- build: avoid significant down time when rebuilding between changes 2025-12-06
- feat: convert gregorian date to quintus in plain text upon request 2025-12-05
- fix: keep the current date selected if changing timezone same date 2025-12-02
- feat: show a date for the timezone sent in requests to time server 2025-12-01
- feat: include metadata meant to find these pages from between site 2025-11-17
- feat: link to a checkout page for selling calendars of soon months 2025-11-12
- fix: inline radio buttons on mobile devices for a dense conversion 2025-10-16
- feat: convert dates with a checkbox selection upon quintus moments 2025-10-14
- build: package compiled css for faster serving on the known domain 2025-10-10
- feat: serve web pages over the secure connection using certificate 2025-10-08
- build: update dependencies to most recent released version changes 2025-10-07
- build: replace additional providers from infrastructure registries 2025-10-07
- chore: update the license year at the turn of the quintus calendar 2025-06-01
- feat: scroll back into the past or toward future dates in infinite 2025-06-01
- test: compare the returned quintus epoch from current moments time 2025-06-01
- build: package dependencies before tullius setups as configuration 2025-05-26
- feat: print quintus dates in a calendar matching related gregorian 2025-05-26
- fix!: keep a matching leap year when converting the current moment 2024-12-02
- fix!: start the new year during nil period but align month january 2024-11-30
- fix: open the http port for web access without using a certificate 2024-10-12
- feat: write the current time and comparison to a webpage on domain 2024-09-28
- fix!: follow the quintus calendar when handling the time protocols 2024-09-26
- feat: await incoming requests on a stable public quintus sh domain 2024-09-26
- feat: serve incoming time requests with configured cloud computing 2024-09-26
- ci: confirm changes are made to the changelog before merges happen 2024-09-17
- ci: confirm udp ntp requests for timings pass testing with success 2024-09-17
- feat: serve ntp times over udp connections using the gregorian utc 2024-09-16

[calver]: https://calver.org
[commits]: https://www.conventionalcommits.org/en/v1.0.0/
[quintus]: https://quintus.sh/now
