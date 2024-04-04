package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// Read input from user
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Send message to server")
		fmt.Println("2. View message history")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter message: ")
			scanner.Scan()
			message := scanner.Text()

			// Send message to server
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error sending message:", err.Error())
				break
			}

			// Receive response from server
			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error receiving response:", err.Error())
				break
			}
			fmt.Println("Server response:", response)
		case "2":
			// Request message history from server
			_, err := conn.Write([]byte("history\n"))
			if err != nil {
				fmt.Println("Error requesting history:", err.Error())
				break
			}

			// Receive and print message history from server
			history, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error receiving history:", err.Error())
				break
			}
			fmt.Println("Message History:")
			fmt.Println(strings.TrimSpace(history))
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}
