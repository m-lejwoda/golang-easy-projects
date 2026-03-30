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
		"HTTP/1.0 200 OK\r\n"+
			"Date: Wed, 25 Mar 2026 20:00:00 GMT\r\n"+
			"Server: MyCustomServer/1.0\r\n"+
			"Content-Type: text/html\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n"+
			"%s", len(body), body)
	return simple_request
}
func sendImage() string {
	path := filepath.Join("/home/michal/PyCharmMiscProject/golang-easy-projects", "village.jpg")
	data, err := os.ReadFile(path)
	fmt.Println("path", path)
	fmt.Println(string(data))
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	return "test"
}

func createConnection() {
	fmt.Println("createConnection")
	// var simple_request = sendSimpleHttpRequest()
	var simple_request = sendImage()
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		//handle error
	}
	print(simple_request)
	print(conn)
	// conn.Write([]byte(simple_request))
}

func main() {
	fmt.Println("Client")
	createConnection()
}
