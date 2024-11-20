package websockets

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader2 = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients8  = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
	mu        sync.Mutex
)

func WsChat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader2.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients8[conn] = true
	mu.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients8, conn)
			mu.Unlock()
			break
		}
		broadcast <- string(msg)
	}
}

func HandleChatMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				client.Close()
				delete(clients8, client)
			}
		}
		mu.Unlock()
	}
}
