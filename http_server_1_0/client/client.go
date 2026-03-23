package main

import (
	"fmt"
	"net"
)

func createConnection() {
	fmt.Println("createConnection")
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		//handle error
	}
	fmt.Println(conn, "GET / HTTP/1.0\r\n\r\n")
	message := "Moja bardzo długa wiadomość, której rozmiaru nie znam..."
	n, err := conn.Write([]byte(message))
	fmt.Println("n", n)
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// fmt.Println("status", status)
}

func main() {
	fmt.Println("Client")
	createConnection()
}
