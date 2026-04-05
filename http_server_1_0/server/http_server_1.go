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

var AvailableContentTypes = []string{"text/html", "image/jpeg"}

func processHtml(*bufio.Scanner) {
	fmt.Println("Process Html")
}

func bytesToImage(data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	img, format, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	fmt.Println("Zdekodowano format:", format)
	return img, nil
}

func processContentTypes(scanner *bufio.Scanner, msg HTTPMessage) {
	if msg.Headers["Content-Type"] == "text/html" {
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			msg.Body += line
		}
	} else if msg.Headers["Content-Type"] == "image/jpeg" {
		fmt.Println("Image")
		var data []byte

		for scanner.Scan() {
			line := scanner.Bytes()
			data = append(data, line...)
		}
		err := os.WriteFile("../../upload/village.jpg", data, 0666)
		if err != nil {
			fmt.Println("Can't save file", err)
		}
	}
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

	processBody(reader, &msg)

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
		fmt.Println("Koniec połączenia lub błąd:", err)
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

func processBody(reader io.Reader, msg *HTTPMessage) {
	lenStr, ok := msg.Headers["Content-Length"]
	if !ok {
		return
	}

	length, _ := strconv.Atoi(lenStr)
	data := make([]byte, length)

	_, err := io.ReadFull(reader, data)
	if err != nil {
		fmt.Println("Błąd czytania danych:", err)
		return
	}

	if msg.Headers["Content-Type"] == "image/jpeg" {
		os.MkdirAll("./upload", 0755)
		os.WriteFile("./upload/village.jpg", data, 0644)
		fmt.Println("Obrazek zapisany!")
	} else {
		msg.Body = string(data)
	}
}
