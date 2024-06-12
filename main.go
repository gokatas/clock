// Tcpserver is a network clock that periodically writes the time. Adapted from
// https://github.com/adonovan/gopl.io/tree/master/ch8/clock2.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:1155")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil { // e.g., connection aborted
			log.Print(err)
			continue
		}
		go handle(conn) // handle connections concurrently
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil { // e.g., client disconnected
			log.Print(err)
			return
		}
		time.Sleep(time.Second)
	}
}
