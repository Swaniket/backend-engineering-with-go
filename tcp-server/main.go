package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			// handle error
		}

		// Goroutine handle multiple connections
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Create a new reader from the connection
	reader := bufio.NewReader(conn)

	// Read the command line from client
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(conn, "Error reading command %v\n", err)
		return
	}

	// Trim the newline char and split the line into command and resource
	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)

	if len(parts) != 2 {
		fmt.Fprint(conn, "Invalid command format. Expected format: COMMAND:RESOURCE\n")
		return
	}

	command := parts[0]
	resource := parts[1]

	log.Printf("Revieved command: %s %s \n", command, resource)

	// Handle the command & ROUTING
	switch command {
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintf(conn, "Unknown command %s\n", command)
	}
}

func handleGet(conn net.Conn, resource string) {
	// Implement GET logic here. for example: fetching a user by ID from a database
	fmt.Fprintf(conn, "Get command recieved for resource: %s\n", resource)
}
