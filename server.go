package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

// Client holds information about a connected client
type Client struct {
	Conn net.Conn
	Name string
}

// Message struct to hold the message information
type Message struct {
	From    string
	To      string
	Content string
}

var clients = make(map[string]Client) // Concurrent map to hold active clients
var mutex = &sync.Mutex{}             // Mutex for handling concurrent access to the map

func main() {
	fmt.Println("Enter port number:")
	var port string
	fmt.Scanln(&port)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer listener.Close()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// New client connected, spawn a goroutine to handle communication
			go handleClient(conn)
		}
	}()

	// Server routine for handling user exit command
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToUpper(input) == "EXIT" {
			for _, client := range clients {
				err := gob.NewEncoder(client.Conn).Encode(Message{From: "Server", Content: "Server stopped"})
				if err != nil {
					fmt.Println(err.Error())
				}
				client.Conn.Close()
			}
			os.Exit(0)
		}
	}
}

// handleClient deals with incoming messages from each client
func handleClient(conn net.Conn) {
	var client Client

	// Get client username
	err := gob.NewDecoder(conn).Decode(&client.Name)
	if err != nil {
		fmt.Println(err.Error())
		conn.Close()
		return
	}

	client.Conn = conn
	mutex.Lock()
	clients[client.Name] = client
	mutex.Unlock()

	// Client connected, wait for messages
	for {
		var msg Message
		err := gob.NewDecoder(conn).Decode(&msg)
		if err != nil {
			mutex.Lock()
			delete(clients, client.Name)
			mutex.Unlock()
			fmt.Println(err.Error())
			conn.Close()
			return
		}

		if msg.To != "" {
			mutex.Lock()
			toClient, ok := clients[msg.To]
			mutex.Unlock()
			if ok {
				err = gob.NewEncoder(toClient.Conn).Encode(msg)
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				err = gob.NewEncoder(conn).Encode(Message{From: "Server", Content: "User not found"})
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}
}
