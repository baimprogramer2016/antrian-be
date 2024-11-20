package websockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader4 = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

func WsSample(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader4.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("New WebSocket connection")

	for {
		// Read message from client
		msgType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		// Print message
		log.Printf("Received: %s\n", p)
		// Send the same message back to the client
		if err := conn.WriteMessage(msgType, p); err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Hello World"))
}
