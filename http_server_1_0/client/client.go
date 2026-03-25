package main

import (
	"fmt"
	"net"
)

func createConnection() {
	fmt.Println("createConnection")
	var simple_request = `HTTP/1.0 200 OK
		Date: Wed, 25 Mar 2026 20:00:00 GMT
		Server: MyCustomServer/1.0
		Content-Type: text/html
		Content-Length: 45

		<html>
		<body>Witaj świecie!</body>
		</html>
		`
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		//handle error
	}
	fmt.Println(conn, "GET / HTTP/1.0\r\n\r\n")
	// message := "Moja bardzo długa wiadomość, której rozmiaru nie znam..."
	n, err := conn.Write([]byte(simple_request))
	fmt.Println("n", n)
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// fmt.Println("status", status)
}

func main() {
	fmt.Println("Client")
	createConnection()
}
