package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func sendSimpleHttpRequest() string {
	var body = `<html>
					<body>Witaj świecie!</body>
				</html>`
	var simple_request = fmt.Sprintf(
		"GET / HTTP/1.0\r\n"+
		"Host: localhost\r\n"+
		"User-Agent: MyGoClient/1.0\r\n"+
			"Date: Wed, 25 Mar 2026 20:00:00 GMT\r\n"+
			// "Server: MyCustomServer/1.0\r\n"+
			"Content-Type: text/html\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n"+
			"%s", len(body), body)
	return simple_request
}
func sendImage() []byte {
	path := filepath.Join("/home/michal/PyCharmMiscProject/golang-easy-projects", "village.jpg")
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Problem with read")
	}
	var headers = fmt.Sprintf(
		"POST /upload HTTP/1.0\r\n"+
			"Date: Wed, 25 Mar 2026 20:00:00 GMT\r\n"+
			"Server: MyCustomServer/1.0\r\n"+
			"Content-Type: image/jpeg\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n",
		len(data))

	fullResponse := append([]byte(headers), data...)
	return fullResponse
}

func createConnection() {
	fmt.Println("createConnection")
	// var simple_request = sendSimpleHttpRequest()
	var simple_request = sendImage()
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Connection error:", err)
        return
	}

	_, err = conn.Write([]byte(simple_request))
	if err != nil {
        fmt.Println("Błąd zapisu:", err)
        return
    }
	if tcpConn, ok := conn.(*net.TCPConn); ok {
        tcpConn.CloseWrite()
    }
	buffer := make([]byte, 4096)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Błąd podczas czytania odpowiedzi:", err)
        return
    }
	fmt.Printf("Otrzymano odpowiedź (%d bajtów):\n%s\n", n, string(buffer[:n]))
}

func main() {
	fmt.Println("Client")
	createConnection()
}
