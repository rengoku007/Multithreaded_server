package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go do(conn) // i can add the do part in a sepeerate thread main thread will go to top of for loop very quickly and do will run seperate with other thread
	}
}

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(8 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n  Hello world\r\n"))
	conn.Close()
}
