package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

type HTTPMessage struct {
	StartLine string
	Headers   map[string]string
	Body      string
}

func parseHttp(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	msg := HTTPMessage{
		Headers: make(map[string]string),
	}
	if scanner.Scan() {
		msg.StartLine = scanner.Text()
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		splitted := strings.SplitN(line, ":", 2)
		msg.Headers[splitted[0]] = splitted[1]
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("test")

			break
		}
		msg.Body += line
		fmt.Printf("%+v\n", msg)
	}

}

func removeConnection(conn net.Conn) {
	fmt.Println("Remove Connection", conn)
}

func handleConnection(conn net.Conn) {
	fmt.Println("Added Connection", conn)
	reader := bufio.NewReader(conn)
	for {
		parseHttp(reader)

		// if message != EOF
		// splitRequest(message)

	}

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
