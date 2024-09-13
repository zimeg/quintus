package udp

import (
	"fmt"
	"log"
	"net"
)

// UDP implements a UDP connection with some helping functions
type UDP struct {
	*net.UDPConn
	packetsize int
}

// Options can be used to configure the UDP connection
type Options struct {
	Port       int
	PacketSize int
}

// Start begins a UDP connection with provided options
func Start(opts Options) (UDP, []byte, error) {
	port := fmt.Sprintf(":%d", opts.Port)
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return UDP{}, []byte{}, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return UDP{}, []byte{}, err
	}
	log.Printf("UDP server listening for NTP requests on port %s", port)
	udp := UDP{
		conn,
		opts.PacketSize,
	}
	buff := make([]byte, udp.packetsize)
	return udp, buff, nil
}

// Read gathers the requester address from incoming connections
func (conn UDP) Read(b []byte) (*net.UDPAddr, error) {
	n, addr, err := conn.ReadFromUDP(b)
	if err != nil {
		return nil, err
	}
	if n != conn.packetsize {
		return nil, fmt.Errorf("Missing bytes in UDP request (%d/%d)", n, conn.packetsize)
	}
	return addr, nil
}

// Write outputs the response to the requesting address
func (conn UDP) Write(b []byte, addr *net.UDPAddr) error {
	n, err := conn.WriteToUDP(b, addr)
	if err != nil {
		return err
	}
	if n != conn.packetsize {
		return fmt.Errorf("Missing bytes in UDP response (%d/%d)", n, conn.packetsize)
	}
	return nil
}
