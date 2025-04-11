package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Define the server URL
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Printf("Connecting to %s\n", u.String())

	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Handle Ctrl+C gracefully
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Start a goroutine to read messages
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Send a message every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt received, closing connection...")

			// Send a close message and wait for the server to respond
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Close message error:", err)
				return
			}
			time.Sleep(time.Second)
			return
		}
	}
}
