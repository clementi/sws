package main

import (
	"fmt"
	"net"
	"os"
)

const (
	host            = ""
	port            = "8080"
	reqBufSizeBytes = 8192
)

func main() {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Listening on " + host + ":" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) error {
	requestBuf := make([]byte, reqBufSizeBytes)

	if _, err := conn.Read(requestBuf); err != nil {
		return err
	}

	response := `HTTP/1.1 200 OK
Content-Type: text/plain

Hello, World!`

	if _, err := conn.Write([]byte(response)); err != nil {
		return err
	}

	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
