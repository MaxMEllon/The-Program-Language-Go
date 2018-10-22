package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	timesByZone := make(chan string)

	for _, arg := range os.Args[1:] {
		query := strings.Split(arg, "=")
		zone, address := query[0], query[1]
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go putTimes(conn, zone, timesByZone)
	}

	for time := range timesByZone {
		fmt.Print(time)
	}
}

func putTimes(conn net.Conn, zone string, timesByZone chan<- string) {
	stream := bufio.NewScanner(conn)
	for stream.Scan() {
		timesByZone <- fmt.Sprintf("[%10s]:\t%s\n", zone, stream.Text())
	}
}
