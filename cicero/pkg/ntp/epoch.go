package ntp

import (
	"time"
)

// Epoch returns the seconds since the NTP epoch 1900-01-01 00:00:00
//
// The seconds between NTP and Unix epochs are added as the offset
func Epoch(t time.Time) uint64 {
	return uint64(2208988800+t.Unix()) << 32
}
