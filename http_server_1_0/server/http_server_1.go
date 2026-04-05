package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func responseMessage(status int, body string) []byte {
    length := len(body)
    var headers = fmt.Sprintf("HTTP/1.0 %d OK\r\n"+
        "Date: Sun, 05 Apr 2026 12:00:00 GMT\r\n"+
        "Server: Apache/1.3.0\r\n"+
        "Content-Type: text/html\r\n"+
        "Content-Length: %d\r\n"+
        "\r\n",
        status, length)
    
    fullResponse := append([]byte(headers), []byte(body)...)
    return fullResponse
}

func bytesToImage(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	fmt.Println("Format decoded:", format)
	return img, nil
}


type HTTPMessage struct {
	StartLine string
	Headers   map[string]string
	Body      string
}

func parseHttp(reader *bufio.Reader, conn net.Conn) error {
	msg := HTTPMessage{
		Headers: make(map[string]string),
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	msg.StartLine = strings.TrimSpace(line)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || strings.TrimSpace(line) == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			msg.Headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	processBody(reader, &msg, conn)

	return nil
}
func removeConnection(conn net.Conn) {
	fmt.Println("Remove Connection", conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Added Connection", conn)

	reader := bufio.NewReader(conn)

	if err := parseHttp(reader, conn); err != nil {
		fmt.Println("End of connection or error:", err)
	}
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

func processBody(reader io.Reader, msg *HTTPMessage, conn net.Conn) {
	lenStr, ok := msg.Headers["Content-Length"]
	if !ok {
		return
	}

	length, _ := strconv.Atoi(lenStr)
	data := make([]byte, length)

	_, err := io.ReadFull(reader, data)
	if err != nil {
		fmt.Println("Problem with data reading:", err)
		conn.Write(responseMessage(400, "Something is wrong with request"))
		return
	}

	if msg.Headers["Content-Type"] == "image/jpeg" {
		os.MkdirAll("./upload", 0755)
		err = os.WriteFile("./upload/village1.jpg", data, 0644)
		if err != nil{
			conn.Write(responseMessage(400, fmt.Sprintf("There was a problem with data send: %s", err)))
		}
		conn.Write(responseMessage(201, "Congratulation. You have sent image"))
	} else {
		msg.Body = string(data)
	}
}
