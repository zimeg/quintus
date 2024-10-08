# Changelog

All changes that happen to this project will be documented in this file.

Updates follow a [conventional commit][commits] style and the project is
versioned with [calendar versioning][calver] using the [quintus][quintus]
calendar.

## Changes

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
