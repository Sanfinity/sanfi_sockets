package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// ============================
// Step 1: Create a WebSocket Upgrader
// ============================
// This helps "upgrade" an HTTP request to a WebSocket connection.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections (for demo purposes)
		return true
	},
}

// ============================
// Step 2: Handle WebSocket Requests
// ============================
// This function will be called whenever someone accesses /ws
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the incoming HTTP request to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close() // Clean up when done

	log.Println("Client connected")

	// Echo loop: Read a message and send it back
	for {
		// Read a message from the client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received: %s", message)

		// Send the same message back to the client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

// ============================
// Step 3: Start the HTTP Server
// ============================
// This serves a basic HTTP page and handles the /ws WebSocket route
func main() {
	// Serve a welcome page at "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<html>
				<body>
					<h1>Hello!</h1>
					<p>This is an HTTP server. Connect to <code>/ws</code> for WebSocket.</p>
				</body>
			</html>
		`)
	})

	// Handle WebSocket connections at "/ws"
	http.HandleFunc("/ws", wsHandler)

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
