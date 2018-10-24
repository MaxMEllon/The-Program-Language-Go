package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	tcp, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("is not tcp conn")
	}
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcp)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	tcp.CloseWrite()
	<-done
	tcp.CloseRead() // 入力後閉じても最後までreverbから送られた内容を出力する
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
