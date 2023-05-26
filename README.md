# A simple chat room Application in Go

This a simple chat room application that supports only private message. It is a simple but alsoeffective demonstration of networking and concurrency in Go. The application consists of two main parts - a server and a client.

## Design

The application is designed with two main components:

1. **Server**: Handles incoming connections from clients, routes messages between clients.

2. **Client**: Connects to the server, sends messages to other users through the server, and receives messages from other users.

Each client runs in its own process. The server can handle multiple concurrent connections from clients and processes each client in its own goroutine.

## Data Flow

1. **Client**:

   - Asks for the server's host address, port number, and a username.
   - Tries to connect to the server.
   - On successful connection, starts a goroutine to listen for incoming messages.
   - Sends user input to the server in the format 'To:Message'. User input is read from the command line.

2. **Server**:
<<<<<<< HEAD

=======
>>>>>>> 433436ae768d31364ef506afb23e06a192433264
   - Accepts incoming connections from clients and spawns a goroutine for each.
   - Each goroutine listens for incoming messages from its client.
   - When a message is received, it's routed to the intended recipient, if they're connected. Otherwise, the sender is notified that the user is not found.

## Running the Application

### Server

1. Open a terminal window.
<<<<<<< HEAD
2. Navigate to the directory to the project root.
3. Run the command `go run ./server`.
=======
2. Navigate to the directory containing the `server.go` file.
3. Run the command `go run server.go`.
>>>>>>> 433436ae768d31364ef506afb23e06a192433264
4. Enter the desired port number for the server to listen on.

### Client

1. Open a new terminal window (separate from the server).
<<<<<<< HEAD
2. Navigate to the directory to the project root.
3. Run the command `go run ./client`.
=======
2. Navigate to the directory containing the `client.go` file.
3. Run the command `go run client.go`.
>>>>>>> 433436ae768d31364ef506afb23e06a192433264
4. Enter the server's host address and port number, and a username for this client.
5. Once connected, you can send messages in the format 'To:Message'.
