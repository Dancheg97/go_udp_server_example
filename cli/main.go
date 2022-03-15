package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	p := make([]byte, 8192)
	conn, err := net.Dial("udp", "127.0.0.1:4271")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hella, hello")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
