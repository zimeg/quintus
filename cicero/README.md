# ⏰cicero

Here within are **handlers** that receive incoming requests for timings that
match a calendar.

```sh
$ make start      # Start the server
$ sntp localhost  # Ask for the time
```

## Serving time with a time server

The [Network Time Protocol][ntp] (NTP) has a specification for how timed
responses should be sent - [RFC5905][rfc]. This is implemented to respond to
incoming requests for the time using a UDP connection on port `123`.

[rfc]: https://datatracker.ietf.org/doc/html/rfc5905
[ntp]: https://en.wikipedia.org/wiki/Network_Time_Protocol
