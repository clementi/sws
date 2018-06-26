package main

import (
	"fmt"
	"net"
)

const host = ""
const port = ":8080"

func main() {
	if listener, err := net.Listen("tcp", port); err != nil {
		fmt.Println(err.Error())
	} else {
		defer listener.Close()
		for {
			if conn, err := listener.Accept(); err != nil {
				fmt.Println(err.Error())
			} else {
				go handleConnection(conn)
			}
		}
	}
}

func handleConnection(conn net.Conn) error {
	requestData := make([]byte, 1024)

	if _, err := conn.Read(requestData); err != nil {
		return err
	}

	response := `
HTTP/1.1 200 OK
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
