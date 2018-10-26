// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

const duration = time.Minute * 5

func parseUser(msg string) (string, bool) {
	matches := regexp.MustCompile(`\[(.+)\]`).FindAllStringSubmatch(msg, -1)
	if matches == nil {
		return "", false
	}
	return matches[0][1], true
}

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)

	who := conn.RemoteAddr().String()
	if input.Scan() {
		firstMsg := input.Text()
		username, success := parseUser(firstMsg)
		if success {
			who = username
		}

	}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	// 入室後放置も切断するように
	timer := time.AfterFunc(duration, func() {
		conn.Close()
	})
	for input.Scan() {
		timer.Stop()
		messages <- who + ": " + input.Text()
		timer = time.AfterFunc(duration, func() {
			conn.Close()
		})
	}

	// NOTE: ignoring potential errors from input.Err()
	timer.Stop()
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
