package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// Message struct to hold the message information
type Message struct {
	From    string
	To      string
	Content string
}

func main() {
	//Connecting to server

	fmt.Println("Enter host address:")
	var host string
	fmt.Scanln(&host)

	fmt.Println("Enter port number:")
	var port string
	fmt.Scanln(&port)

	fmt.Println("Enter username:")
	var username string
	fmt.Scanln(&username)

	// Dial server
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Failed to connect, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
		main()
		return
	}

	// Send username to server
	err = gob.NewEncoder(conn).Encode(username)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Goroutine to handle incoming messages
	go func() {
		for {
			var msg Message
			err := gob.NewDecoder(conn).Decode(&msg)
			if err != nil {
				fmt.Println(err.Error())
				conn.Close()
				os.Exit(0)
			}
			fmt.Println(msg.From + ": " + msg.Content)
		}
	}()

	fmt.Println("Connected! Enter message in the format 'To:Message'. Type 'EXIT' to quit.")

	// Main routine for user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			os.Exit(0)
		}
		tokens := strings.SplitN(input, ":", 2)
		if len(tokens) != 2 {
			fmt.Println("Invalid format. Use 'To:Message'")
			continue
		}
		msg := Message{
			From:    username,
			To:      tokens[0],
			Content: tokens[1],
		}
		err := gob.NewEncoder(conn).Encode(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
