package main

import (
	"fmt"
	"net"
)

func removeConnection(conn net.Conn) {
	fmt.Println("Remove Connection", conn)
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	fmt.Println("Added Connection", conn)
	n, err := conn.Read(buffer)
	if err != nil {
		//handle err
	}
	message := string(buffer[:n])
	fmt.Println("message", message)
	removeConnection(conn)
}

func startHttpServer10() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}

}
