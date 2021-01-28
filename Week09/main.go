package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync/atomic"
)

var uid int64 = 1

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	defer listen.Close()
	fmt.Println("tcp server running on 12345:")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	connChan := make(chan string, 16)

	go send(conn, connChan)
	for {
		receive, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("[uid: %v] [addr: %v]: %v", int(atomic.AddInt64(&uid, 1)), conn.RemoteAddr().String(), receive)
		connChan <- receive
	}
}

func send(conn net.Conn, message chan string) {
	for msg := range message {
		conn.Write([]byte(msg))
	}
}
