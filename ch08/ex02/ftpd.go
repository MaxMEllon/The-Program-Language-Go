package main

import (
	"log"
	"net"

	"github.com/maxmellon/The-Program-Language-Go/ch08/ex02/ftpd"
)

func main() {
	listener, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		log.Print("connected")
		if err != nil {
			log.Print(err)
			continue
		}
		go ftpd.HandleConnect(conn)
	}
}
