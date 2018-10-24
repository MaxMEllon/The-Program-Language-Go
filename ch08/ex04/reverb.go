package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

//!+
func echo(c net.Conn, shout string, delay time.Duration, waitGroup sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	waitGroup.Done() // 実行中状態から完了状態に
}

func handleConn(c *net.TCPConn, waitGroup sync.WaitGroup) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		waitGroup.Add(1) // goルーチン 一個分 インクリメント
		go echo(c, input.Text(), 1*time.Second, waitGroup)
	}
	waitGroup.Wait() // サブルーチン delta 個の処理を待つ
	c.CloseWrite()   // 待ってからtcpの書き込み側をclose
}

//!-

func main() {
	var waitGroup sync.WaitGroup
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		tcp, ok := conn.(*net.TCPConn)
		if !ok {
			log.Fatal("is not tcp conn")
		}
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(tcp, waitGroup)
	}
}
