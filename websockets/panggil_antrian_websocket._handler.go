package websockets

import (
	monitorantrianrepository "be-mklinik/repositories/monitor_antrian_repository"
	"be-mklinik/requests"
	monitorantrianservice "be-mklinik/services/monitor_antrian_service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

var (
	clients  = make(map[*websocket.Conn]bool)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WsAntrian(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("New client connected ", clients)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client Close: ", msg, err)
			delete(clients, conn)
			break
		}

		log.Println("Client :", string(msg))
	}
}

type antrianHandler struct {
	monitorAntrianService monitorantrianservice.MonitorAntrianService
}

func NewPanggilAntrianHandler(db *gorm.DB) *antrianHandler {
	repository := monitorantrianrepository.NewMonitorAntrianRepository(db)
	monitorAntrianService := monitorantrianservice.NewMonitorAntrianService(repository)
	return &antrianHandler{monitorAntrianService: monitorAntrianService}
}

func (a *antrianHandler) PanggilAntrian(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var panggilanRequest requests.ParamUpdateAntrianRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&panggilanRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := a.monitorAntrianService.UpdatePanggilanAntrian(panggilanRequest)

	if err != nil {
		http.Error(w, "Gagal Update", http.StatusInternalServerError)
		return
	}

	fmt.Println(result)

	defer r.Body.Close()

	monitorAntrianService, err := a.monitorAntrianService.GetAllMSeqnoAntrianByDay()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for client := range clients {

		json, err := json.Marshal(monitorAntrianService)
		if err != nil {
			log.Println("Error broadcasting message:", err)
			continue
		}
		err = client.WriteMessage(websocket.TextMessage, []byte(json))
		if err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(clients, client)
		}
	}

}
