package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/Bhanuteja002/A-simple-chat-room-Application-in-Go/message" // update this to the actual path if necessary
)

const maxRetries = 10

func main() {
	fmt.Println("Enter host address:")
	var host string
	fmt.Scanln(&host)

	fmt.Println("Enter port number:")
	var port string
	fmt.Scanln(&port)

	fmt.Println("Enter username:")
	var username string
	fmt.Scanln(&username)

	// Keep dialing
	for retries := 0; retries < maxRetries; retries++ {
		conn, err := net.Dial("tcp", host+":"+port)
		if err == nil {
			err = gob.NewEncoder(conn).Encode(username)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// Goroutine to handle incoming messages
			go func() {
				for {
					var msg message.Message
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

				msg := message.Message{
					From:    username,
					To:      tokens[0],
					Content: tokens[1],
				}
				err := gob.NewEncoder(conn).Encode(msg)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			break
		}

		fmt.Printf("Failed to connect, retrying in 5 seconds... (Attempt %d of %d)\n", retries+1, maxRetries)
		time.Sleep(5 * time.Second)

		if retries == maxRetries-1 {
			fmt.Println("Max retries exceeded, quitting...")
			os.Exit(1)
		}
	}
}
