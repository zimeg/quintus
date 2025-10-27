package ntp

import (
	"encoding/binary"

	"github.com/zimeg/quintus/cicero/pkg/now"
	"github.com/zimeg/quintus/cicero/pkg/utc"
)

// NTPPacket contains details about the values of an NTP packet
//
// https://datatracker.ietf.org/doc/html/rfc5905
type NTPPacket struct {
	LiVnMode   uint8  // Leap indicator, version, mode
	Stratum    uint8  // Stratum level
	Poll       int8   // Polling interval
	Precision  int8   // Precision
	RootDelay  uint32 // Root delay
	RootDisp   uint32 // Root dispersion
	RefID      uint32 // Reference identifier
	RefTime    uint64 // Reference time          (current moment)
	OrigTime   uint64 // Origin timestamp        (client request)
	RcvTime    uint64 // Receive timestamp       (server started)
	TxTime     uint64 // Transmit timestamp      (responses sent)
	packetsize int
}

// New creates an NTP packet in response to the incoming request time
func New(request []byte) NTPPacket {
	original := binary.BigEndian.Uint64(request[40:])
	moment := now.Moment(utc.Current().ToTime())
	packet := NTPPacket{
		LiVnMode:   0x1C, // LI:0 (no warning), VN:4 (version 4), Mode:4 (server)
		Stratum:    1,    // Primary server
		Poll:       4,
		Precision:  -20,
		RefID:      0x58495645, // Experiment "XIVE" should start with "X" for now
		RefTime:    moment.Offset(),
		OrigTime:   original,
		RcvTime:    moment.Offset(),
		TxTime:     moment.Offset(),
		packetsize: len(request),
	}
	return packet
}

// Marshal formats the packet into a response value
func (p *NTPPacket) Marshal() []byte {
	b := make([]byte, p.packetsize)
	b[0] = p.LiVnMode
	b[1] = p.Stratum
	b[2] = byte(p.Poll)
	b[3] = byte(p.Precision)
	binary.BigEndian.PutUint32(b[4:], p.RootDelay)
	binary.BigEndian.PutUint32(b[8:], p.RootDisp)
	binary.BigEndian.PutUint32(b[12:], p.RefID)
	binary.BigEndian.PutUint64(b[16:], p.RefTime)
	binary.BigEndian.PutUint64(b[24:], p.OrigTime)
	binary.BigEndian.PutUint64(b[32:], p.RcvTime)
	binary.BigEndian.PutUint64(b[40:], p.TxTime)
	return b
}
