# Changelog

All changes that happen to this project will be documented in this file.

Updates follow a [conventional commit][commits] style and the project is
versioned with [calendar versioning][calver] using the [quintus][quintus]
calendar.

## Changes

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
[quintus]: https://api.o526.net/v1/calendar/today
