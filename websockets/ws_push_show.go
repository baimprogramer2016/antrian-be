package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients3 = make(map[*websocket.Conn]bool)
var upgrader3 = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler untuk koneksi WebSocket
func ListenCall(w http.ResponseWriter, r *http.Request) {
	// Upgrade koneksi HTTP ke WebSocket
	conn, err := upgrader3.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Tambahkan koneksi baru ke map
	clients3[conn] = true
	log.Println("New client connected")

	// Terus tunggu pesan dari client, tapi tidak melakukan broadcast terus-menerus
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(clients3, conn)
			break
		}
	}
}

// Fungsi untuk memanggil antrian dan melakukan broadcast

type AntrianRequest struct {
	NomorAntrian string `json:"nomor_antrian"`
}

func CallAntrian(w http.ResponseWriter, r *http.Request) {

	var request AntrianRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Nomor antrian yang diterima dari input (misalnya: {"queueNumber": "Antrian #5"})
	queueNumber := request.NomorAntrian
	if queueNumber == "" {
		// Jika tidak ada nomor antrian dalam request, buat nomor antrian acak
		queueNumber = fmt.Sprintf("Antrian #%d", 1+len(clients3))
	}

	fmt.Println(queueNumber)
	// Broadcast pesan ke semua client yang terhubung
	for client := range clients3 {
		err := client.WriteMessage(websocket.TextMessage, []byte(queueNumber))
		if err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(clients3, client)
		}
	}
	fmt.Fprintf(w, "Antrian dipanggil: %s", queueNumber)
}

func HandleCallAntrian() {

}
