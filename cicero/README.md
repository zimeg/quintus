# ⏰cicero

Here within are **handlers** that receive incoming requests for timings that
match a calendar.

## Serving time with a time server

The [Network Time Protocol][ntp] (NTP) has a specification for how timed
responses should be sent - [RFC5905][rfc]. This is implemented to respond to
incoming requests for the time using a UDP connection on port `123`.

### Starting the server

Future moments can be realized on request with a started `go` server:

```sh
$ make start
...
2024/10/08 17:41:20 UDP server listening for NTP requests on port :123
2024/10/08 17:41:20 TCP server handling the HTTP requests on port :80
```

### Asking for the time

A request for the current moment can be made with the `sntp` command:

```sh
$ sntp localhost
...
2024-10-08 17:41:24.641235 (+0400) +345599.359352 +/- 230399.572916 localhost ::1 s1 no-leap
```

[rfc]: https://datatracker.ietf.org/doc/html/rfc5905
[ntp]: https://en.wikipedia.org/wiki/Network_Time_Protocol
