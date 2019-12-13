package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	// accept arguments and show usage
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <role> <args>")
		fmt.Println("client serverIP:serverPort msgToSend")
		fmt.Println("server")
		os.Exit(1)
	}

	role := os.Args[1]

	if role == "client" {
		tcpClient(os.Args[2], os.Args[3])
	}
	if role == "server" {
		tcpServer()
	}
}

func tcpServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		fmt.Println("[DEBUG] Waiting on port 8080")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("[DEBUG] ***Connection Established with " + conn.RemoteAddr().String() + "***")
		// echo something back to the client
		conn.Write([]byte("You contacted the server!\nThe time is " + time.Now().String() + "\n"))
		conn.Close()
		fmt.Println("[DEBUG] ***Connection closed***")
	}
}

func tcpClient(serverAddr string, msg string) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	conn.Write([]byte(msg))
}
