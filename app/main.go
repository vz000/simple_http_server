package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
)


type StatusLine struct {
	HttpVersion, StatusCode, ReasonPhrase string
}

/*
type Headers struct {
	todo string
}*/

func newResponse(statusCode int, reasonPhrase string) *StatusLine {
	statusLine := StatusLine{HttpVersion: "HTTP/1.1\n"}
	statusLine.StatusCode = strconv.Itoa(statusCode) + "\n"
	statusLine.ReasonPhrase = reasonPhrase + "\n\r\n" // reason phrase + new line + \r\n
	return &statusLine
}

func client_connection(c net.Conn) {
	dataBuffer := make([]byte, 4096) // big buffer
	var clientData strings.Builder
	var response strings.Builder
	n, err := c.Read(dataBuffer)
	if err != nil {
		fmt.Println("Error reading client data: ", err.Error())
		os.Exit(1)
	}
	clientData.Write(dataBuffer[:n])
	fmt.Println(clientData.String())

	// This will be changed
	// status line 構成
	statusLine := newResponse(200, "OK")
	response.WriteString(statusLine.HttpVersion)
	response.WriteString(statusLine.StatusCode)
	response.WriteString(statusLine.ReasonPhrase)

	// Headers
	response.WriteString("\r\n")
	n, err = c.Write([]byte(response.String()))
	if err != nil {
		fmt.Println("Error sending response: ", err.Error())
		os.Exit(1)
	}
	fmt.Println(response.String())
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
	 	os.Exit(1)
	}

	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	for {
		go client_connection(conn)
	}
}
