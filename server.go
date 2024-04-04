package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var messageHistory []string

func main() {
	// Setting up the server on port 9090
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Server is listening on port 9090")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New client connected:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// Read message from client
		message := scanner.Text()
		fmt.Println("Received message from client:", message)

		// Log message in history
		messageHistory = append(messageHistory, message)

		// Echo message back to client
		conn.Write([]byte("Server: " + message + "\n"))
	}

	fmt.Println("Client disconnected:", conn.RemoteAddr())
}
