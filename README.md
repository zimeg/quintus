# 🗓️ quintus

> five day calendar

## Counting the weeks to make a month

The `quintus` calendar is based on a **five day** week. Connecting six weeks
together makes twelve equal **thirty day** months with just a few days leftover:

| week |   a |   b |   c |   d |   e |
| ---: | --: | --: | --: | --: | --: |
|  `1` |   1 |   2 |   3 |   4 |   5 |
|  `2` |   6 |   7 |   8 |   9 |  10 |
|  `3` |  11 |  12 |  13 |  14 |  15 |
|  `4` |  16 |  17 |  18 |  19 |  20 |
|  `5` |  21 |  22 |  23 |  24 |  25 |
|  `6` |  26 |  27 |  28 |  29 |  30 |

Nuance of the calendar are outlined in [post][post] and elsewhere but this
project focuses on timed implementation.

## Following the Quintus Time Server (QTS)

Part of this project is dedicated to serving the true times in computer format
with the [Network Time Protocol][ntp]:

```sh
$ sntp quintus.sh
...
2024-09-22 22:40:36.017959 (+0700) +345599.982155 +/- 230399.988118 quintus.sh 3.84.124.11 s1 no-leap
```

[post]: https://o526.net/blog/post/five-day-week/
[ntp]: https://en.wikipedia.org/wiki/Network_Time_Protocol
