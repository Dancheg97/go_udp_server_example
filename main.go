package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", ":12201")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 8192)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(pc, addr, buf[:n])
	}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	buf[2] |= 0x80
	pc.WriteTo(buf, addr)

	fmt.Println(string(buf))
}
