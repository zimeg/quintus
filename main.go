package main

import (
	"log"
	"net"

	"github.com/zimeg/quintus/pkg/ntp"
	"github.com/zimeg/quintus/pkg/tcp"
	"github.com/zimeg/quintus/pkg/udp"
)

// respond writes the NTP response to the UDP request
func respond(conn udp.UDP, addr *net.UDPAddr, request []byte) {
	packet := ntp.New(request)
	err := conn.Write(packet.Marshal(), addr)
	if err != nil {
		log.Printf("Failed to write the NTP response: %s", err)
		return
	}
}

// main routes requests seeking time
func main() {
	opts := udp.Options{
		Port:       123,
		PacketSize: 48,
	}
	conn, buff, err := udp.Start(opts)
	if err != nil {
		log.Printf("Failed to start a new UDP process: %s", err)
		return
	}
	defer conn.Close()
	go tcp.Listen()
	for {
		addr, err := conn.Read(buff)
		if err != nil {
			log.Printf("Failed to read a UDP request: %s", err)
			continue
		}
		go respond(conn, addr, buff)
	}
}
